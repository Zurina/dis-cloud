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





    
