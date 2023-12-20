# Snap Roles

## Introduction

In the challenging economic landscape of 2023, students are facing unprecedented difficulties securing internships or full-time employment (FTE) opportunities. Unlike previous years, where job offers were plentiful, students now find themselves submitting hundreds of applications without guarantee of even an interview. The competition is fierce, with an overwhelming number of applicants vying for a limited number of positions at top companies. Job postings fill up rapidly, often within hours, making it challenging for students to stay informed and secure timely opportunities. The struggle is not just in securing offers but in obtaining relevant information about companies and their job postings in a timely manner.

Enter Snap Roles, an application designed to address this pressing issue. The program aims to provide instantaneous information about recent graduate and internship opportunities, with a latency as low as 10-15 minutes after a job posting. Leveraging official APIs from companies ensures the accuracy of the data provided.

## Architecture Diagram
![Architecture Diagram](assets/images/arch_diagram.png)

## Work Until Now
As of now, the repository supports job information retrieval for two major companies:
- Microsoft
- Apple

The current implementation covers only one aspect of the architecture, with Swagger serving as a tool to interact with the APIs. Notably, parameters such as recent graduate, intern status, and others are currently hardcoded.

## How It Works
Snap Roles aims to gather job information directly from official company websites, utilizing freely available APIs. The program employs a scheduler that checks APIs every 15 minutes, resulting in an overall latency of 15 minutes. This ensures that students receive timely updates on job postings. While some companies have well-structured APIs providing timestamp information (e.g., Microsoft), others like Apple only return the date of the job posting. To address this, a system has been implemented to store job IDs and compare them with previous scheduler iterations, effectively identifying jobs posted in the last 15 minutes.

For companies with highly authenticated services, like Meta, web scraping becomes the method of choice. The scheduler can be adapted to run on a web scraper to fetch the latest job information.

Certain companies may initially post generic positions for all teams and later assign candidates accordingly. Additionally, individual managers, both from known and unknown companies (where the scheduler is operational), may post job opportunities on platforms like LinkedIn. For such cases, data scraping from LinkedIn posts or the LinkedIn Posts API can be employed, depending on guidelines.

## Prerequisites
Ensure Go is installed, and execute the following commands in the root folder.

## Running
To initialize and run Swagger, use the following command (this command will build, and run application server):
```
make swag_run
```

To clean the solution and remove binaries, run:
```
make clean
```
