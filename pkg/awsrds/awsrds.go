package awsrds

import (
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
)

func parseDBInstance(d *rds.DBInstance) (bool, string, error) {
	if aws.StringValue(d.DBInstanceStatus) == "creating" {
		return true, "", nil
	}

	if aws.StringValue(d.DBInstanceStatus) != "available" {
		return true, "", fmt.Errorf("instance %s not available: %s",
			aws.StringValue(d.DBInstanceIdentifier),
			aws.StringValue(d.DBInstanceStatus))
	}

	return true, fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s",
		aws.StringValue(d.Endpoint.Address),
		aws.Int64Value(d.Endpoint.Port),
		aws.StringValue(d.DBName),
		aws.StringValue(d.MasterUsername),
		aws.StringValue(d.MasterUsername),
	), nil
}

func findPostgresRDS(svc *rds.RDS, instID string) (bool, string, error) {
	result, err := svc.DescribeDBInstances(nil)
	if err != nil {
		return false, "", fmt.Errorf("DescribeDBInstances failed with %s", err)
	}

	for _, d := range result.DBInstances {
		//fmt.Printf("%s\n%#v\n", aws.StringValue(d.DBInstanceIdentifier), d)
		if aws.StringValue(d.DBInstanceIdentifier) == instID {
			return parseDBInstance(d)
		}
	}

	return false, "", nil
}

func createPostgresRDS(svc *rds.RDS, instID string) (bool, string, error) {
	result, err := svc.CreateDBInstance(
		&rds.CreateDBInstanceInput{
			AllocatedStorage:        aws.Int64(20),
			AutoMinorVersionUpgrade: aws.Bool(false),
			BackupRetentionPeriod:   aws.Int64(0),
			DBInstanceClass:         aws.String("db.t2.micro"),
			DBInstanceIdentifier:    &instID,
			DBName:                  aws.String("sqltest"),
			Engine:                  aws.String("postgres"),
			EngineVersion:           aws.String("11.5"),
			MasterUsername:          aws.String("postgres"),
			MasterUserPassword:      aws.String("postgres"),
			PubliclyAccessible:      aws.Bool(false),
			StorageType:             aws.String("gp2"),
		})
	if err != nil {
		return false, "", fmt.Errorf("CreateDBInstance failed with %s", err)
	}

	return parseDBInstance(result.DBInstance)
}

func EnsurePostgresRDS(instID string) (string, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")})
	if err != nil {
		return "", fmt.Errorf("NewSession failed with %s\n", err)
	}
	svc := rds.New(sess)

	found, s, err := findPostgresRDS(svc, instID)
	if err != nil {
		return "", err
	}

	if !found {
		found, s, err = createPostgresRDS(svc, instID)
		if err != nil {
			return "", err
		}
		if !found {
			return "", fmt.Errorf("instance %s not found after being created", instID)
		}
	}

	if s != "" {
		return s, nil
	}

	log.Printf("AWS RDS: waiting for %s to be available", instID)
	for s == "" {
		log.Print(".")
		time.Sleep(10 * time.Second)

		found, s, err = findPostgresRDS(svc, instID)
		if err != nil {
			return "", err
		}
		if !found {
			return "", fmt.Errorf("instance %s no longer found", instID)
		}
	}

	return s, nil
}
