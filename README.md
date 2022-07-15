# Distributed Systems & Cloud

### 1. Part one. "Should we move to AWS Linux 2022"?

#### **Summary**

What we're considering here is whether we should change the base image of our AMI, which currently is Ubuntu, to AWS Linux 2022.

#### **Motivation**

The motivation for this consideration is that AWS Linux 2022 might be more efficient and can save cost for us.

#### **Research**

First of all, AWS Linux 2022 is a Linux version custom made by AWS. It has a lot of focus on the fact that it is meant to run in the cloud, and therefore, provides extensive security and a high performance execution environment to develop/run cloud native applications. 

Pros

- Usage of AWS Linux 2022 is available at no additional cost.
- Specifically optimized for AWS EC2 and is heavily integrated with other features of AWS which can make your developer life easier if you want to have your EC2 instances integrated with other AWS tools.
- AWS Linux 2022 requires less hardware resources that enable many users to access it.
- Amazon Linux 2022 can subscribe to AWS support, due to its native integration in Amazon cloud.
- You also get little conveniences like being able to use the update infrastructure through an S3 endpoint so your instances don't need Internet access to do a yum update. This feature is quite nice, so you don't need a NAT gateway for that for example, although there other ways to omit that as well. 

Cons

- The biggest con of choosing AWS Linux 2022 would be vendor lock-in. You're tying yourself to AWS, which for the future of the company might not be ideal. Obviously AWS is not a bad place to be locked in, but it's still an important consequence to consider. 

Conclusion

- My conclusion cannot be concluded in my opinion, as it depends on what kind of company we're talking about. If it's a startup in question, I think I'd avoid vendor lock-in and thereby sticking with Ubuntu OS. If however, we're talking about a big company that is already heavily integrated in AWS, I would say switing to AWS Linux 2022 makes more sense.

### 2. Part two. EC2 and custom AMI

#### Step-by-step guide how to create custom AMI

First install go on the EC2 instance using following commands:

- curl -O https://storage.googleapis.com/golang/go1.17.5.linux-amd64.tar.gz
- tar -xvf go1.17.5.linux-amd64.tar.gz
- sudo mv go /usr/local
- export GOPATH=$HOME/work
- export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
- source ~/.profile

And then:

- Log in to the AWS Management Console, display the EC2 page for your region, then click Instances.
- Choose the instance from which you want to create a custom AMI.
- Click Actions, Click Image and Templates, and then click Create Image.
- Fill in the details.
- Click Create Image.

#### Prepare script that is gonna launch your app

    git clone https://github.com/Zurina/dis-cloud.git
    cd dis-cloud
    go build main.go
    chmod +x main
    ./main

### Docker & Docker Compose

To setup the loadbalancer locally, follow following steps:

- In the root of the project, run:

        docker build -t dis-cloud .

- Change directory to loadbalancer.
- Run:

        docker build -t loadbalancer .

- Run:

        docker-compose up

- You should now be able to hit the loadbalancer on http://localhost:8000

### AWS CLI tool

        aws ec2 run-instances --image-id ami-0149623249da586db --placement AvailabilityZone=us-east-1c --count 1 --instance-type t2.micro --key-name mathias --region us-east-1 --tag-specifications 'ResourceType=instance,Tags=[{Key=Name,Value=mat-ins}]'

        aws ec2 create-volume --region us-east-1 --availability-zone us-east-1c --size 1 --tag-specification 'ResourceType=volume,Tags=[{Key=Name,Value=mat-volume}]'

        VOLUME_ID=$(aws ec2 describe-volumes  --region us-east-1 --query "Volumes[*].{ID:VolumeId,AZ:AvailabilityZone,Size:Size}" --filters "Name=tag:Name,Values=mat-volume" | jq .[].ID | tr -d \")

        INSTANCE_ID=$(aws ec2 describe-instances --region us-east-1 --query "Reservations[*].Instances[].InstanceId" --filter 'Name=tag:Name,Values=mat-ins' --output text)      
        
        aws ec2 attach-volume --region us-east-1 --device /dev/xvdb --instance-id $INSTANCE_ID --volume-id $VOLUME_ID

        SNAPSHOT_ID=$(aws ec2 create-snapshot --region us-east-1 --volume-id $VOLUME_ID --tag-specifications 'ResourceType=snapshot,Tags=[{Key=Name,Value=mat-snap}]' | jq .SnapshotId | tr -d \")

        aws ec2 detach-volume --region us-east-1 --device /dev/xvdb --instance-id $INSTANCE_ID --volume-id $VOLUME_ID

        aws ec2 delete-snapshot --region us-east-1 --snapshot-id $SNAPSHOT_ID

        aws ec2 terminate-instances --region us-east-1 --instance-id $INSTANCE_ID --delete-volumes










    
