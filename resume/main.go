package main

import (
	"fmt"
	"net"
	"log"
)

func main() {

	li, err := net.Listen("tcp", ":80")
	if err != nil {
	log.Fatalln(err)
	}
	defer li.Close()
	for {
		connection, err := li.Accept()
		if err != nil {
		log.Fatalln(err)
		}
	
	go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	body := ` 
	<!DOCTYPE html>
	<html lang="en">
	<head>
    	<meta charset="utf-8">
    	<title>Ranjith's Web</title>
	</head>
	<body style="background-color:powderblue;">
    	<h1> RESUME</h1>


	
                             
	<h2>ROFESSIONAL SUMMARY</h2>
	<ul>
	<li>Having 6 years of experience in handling infrastructure automation and <b>micro services</b>.</li>
	<li>Administrating Linux servers utilizing web servers like nginx, apache, tomcat and continuous integration with Jenkins.</li>
	<li>Containerizing applications using <b>Docker, Docker Compose, Kubernetes</b>.</li>
	<li>Having good knowledge on <b>Go programming</b>.</li>
	<li>Creating infrastructure on <b>AWS</b> Cloud.</li>
	<li>Building CI/CD pipeline for Jenkins integrated with GitHub.</li>
	<li>Writing Ansible playbooks as per the requirement.</li>
	<li>Experienced in server builds and software deployment in Unix/Linux environment.</li>
	<li>Managing the infrastructure with automation tool like Ansible.</li>
	<li>Having good knowledge on writing scripts using Bash Shell for day to day activities.</li>
	<li>Experience in designing, installation, administration, troubleshooting and maintenance of servers.</li>
	<li>Hands on experience in handling services and applications like Samba, DNS Bind, DHCP, SSL, SSH, Apache, Tomcat and load balancing on Linux machine.</li>
	<li>File system management: Managing partitions, Virtual Memory Management & LVM.</li>
	</ul>

	<h2>PROFESSIONAL EXPERIENCE: 1</h2>
	<ul>
	<li>Organization	:	Dell SecureWorks </li> 
	<li>Designation 	:	Analyst </li>
	<li>Duration 		:	27 Months </li>
	</ul>
	<h3>Roles and Responsibilities: </h3> 
	<ul>
	<li>Responsible for migration business applications into micro services using docker, Kubernetes.</li>
	<li>Handling infrastructure issues daily including configuration, building systems, application installs, server maintenance and configuration.</li>
	<li>Working on latest technologies like AWS (EC2, VPC, ELB, Autoscaling, IAM, etc...), Docker, Kubernetes.</li>
	<li>Managing the infrastructure with automation tool like Ansible.</li>
	<li>Building CI/CD pipeline for Jenkins integrated with GitHub.</li>
	<li>Responsible for the implementation, troubleshooting, and maintaining operations of network systems.</li>
	<li>Working closely with Engineering (Dev/QA) and Senior Architects to build a product development.</li>
	<li>Maintaining, monitoring and troubleshooting disk space, volumes, Https, site down, Load Average, free memory, CPU Usage, Processes etc. Alerts generated in monitoring tools.</li>
	<li>Performing Root cause analysis (RCA), Corrective and Preventive actions (CAPA).</li>
	<li>Good Knowledge on writing automation scripts using shell scripting for day to day activities.</li>
	</ul>

	<h2>PROFESSIONAL EXPERIENCE: 2</h2>
	<ul>
	<li>Organization	:	Techwave Consulting </li>                                           
	<li>Designation 	:	System Engineer </li>
	<li>Duration 		:	13 Months </li>
	</ul>
	<h3>Roles and Responsibilities: </h3>
	<ul>
	<li>As part of Devops team we handle & maintain data centers of our client for more than 1000+ physical servers & 2000+ virtual hosts for continuous high availability of infrastructure.
	<li>Hands on experience in creating infrastructure on AWS Cloud.</li>
	<li>Wide knowledge on configuring samba, DNS Bind, DHCP, HTTPS and Apache Web Servers on Linux machines.</li>
	<li>Performing L2 & Partial L3 level full lifecycle management for all alerts and production issues like incident logging, initial troubleshooting, doing root cause analysis, updating process Documentation on Linux Machines.</li>
	<li>Generating weekly reports and sharing with the Management about escalating tickets as per Project SLA’S, SLA Compliance, Ticket Volume and Escalation trends. And escalating tickets as per Project SLA’S.</li>
	<li>Responsible for the implementation, troubleshooting, and maintaining operations of network systems.</li>
	Experience in working with Production, Service Delivery, and Support Infrastructure NOC, Ticketing environments on 365 *24/7 Basis.</li>
	</ul>
	<h2>PROFESSIONAL EXPERIENCE: 3</h2>
	<ul>
	<li>Organization	:	Wipro InfoTech (Partner IDC Technologies)</li>
	<li>Designation 	:	LINUX Administrator</li>
	<li>Duration 		:	14 Months </li>
	</ul>
	<h3>Roles and Responsibilities: </h3>
	<ul>
	<li>Worked as a Linux Administrator in all aspects of the infrastructure, manageability of the service, support and deployment automation for the product Development.</li>
	<li>Supporting Linux servers utilizing web servers like apache.</li>
	<li>Linux Performing L1 and Partial L2 level full lifecycle management for all alerts and production issues like incident logging, initial troubleshooting, doing root cause analysis, updating process documentation on Machines.</li>
	<li>Handling infrastructure issues daily, including configuration, building systems, application installs, server maintenance and configuration.</li>
	<li>Monitoring, troubleshooting the performance related Events like disk space, volumes, Https, site down, Load Average, free memory, cpu usage, Processes etc.</li>
	<li>Responsible for the implementation, troubleshooting, and maintaining operations of network systems.</li>     
	</ul>
	<h2>PROFESSIONAL EXPERIENCE: 4</h2> 
	<ul>
	<li>Organization	:    4 WAY SOLUTIONS</li>                                       
	<li>Designation 	:    System Support Engineer</li>
	<li>Duration 	:    20 Months</li>
	</ul>
	<h3>Roles and Responsibilities:</h3>
	<ul>
	<li>Responsible for assisting Lead Engineer in areas of design, configuration and implementation of IT solutions.</li>
	<li>Handle the tasks of identifying, diagnosing, and resolving hardware and software problems.</li>
	<li>Responsible for the implementation, troubleshooting, and maintaining operations of network systems.</li>
	<li>Maintaining, monitoring and troubleshooting disk space, volumes, Https, site down, Load Average, free memory, Processes etc. Alerts generated in monitoring tools etc.</li>
	<li>Strong administration and troubleshooting skills on Linux/UNIX.</li>
	<li>Provide desktop support including creating images specific to client requirements and deal with issues pertaining to hardware and applications.</li>
	<li>Perform routine network maintenance checks as well as configure and manage printers, copiers, and another miscellaneous network equipment.</li>
	</ul>	
	</body>
	</html>
`
fmt.Fprintf(connection, "%v", body)
 connection.Close()
}

