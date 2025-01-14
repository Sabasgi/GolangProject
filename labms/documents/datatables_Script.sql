
CREATE TABLE Country (
    country_id INT PRIMARY KEY AUTO_INCREMENT,
    country_name VARCHAR(255) NOT NULL,
    country_code VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    gst_percentage INT

);

# State Table
CREATE TABLE State (
    state_id INT PRIMARY KEY AUTO_INCREMENT,
    state_name VARCHAR(255) NOT NULL,
    country_id INT,
    FOREIGN KEY (country_id) REFERENCES Country(country_id),
    created_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255)
);

# City Table
CREATE TABLE City (
    city_id INT PRIMARY KEY AUTO_INCREMENT,
    city_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    state_id INT,
    FOREIGN KEY (state_id) REFERENCES State(state_id),
     created_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255)
);


# Lab Table
CREATE TABLE Lab (
    lab_id INT PRIMARY KEY AUTO_INCREMENT,
    lab_name VARCHAR(255) NOT NULL,
    lab_Code VARCHAR(255) NOT NULL UNIQUE,
    created_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255)

    #address
);

# Branch Table
CREATE TABLE Branch (
    branch_id INT PRIMARY KEY AUTO_INCREMENT,
    branch_name VARCHAR(255),
    lab_id INT,
    FOREIGN KEY (lab_id) REFERENCES Lab(lab_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    address VARCHAR(255),
    branch_code VARCHAR(255) NOT NULL UNIQUE,
    city_id INT,
    FOREIGN KEY (city_id) REFERENCES City(city_id)
);

CREATE TABLE Department (
    department_id INT PRIMARY KEY AUTO_INCREMENT,
    department_name VARCHAR(255) NOT NULL,
    lab_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (lab_id) REFERENCES Lab(lab_id),
    branch_id INT NOT NULL,
    FOREIGN KEY (branch_id) REFERENCES branch(branch_id),
    description TEXT
);

CREATE TABLE Service (
    service_id INT PRIMARY KEY AUTO_INCREMENT,
    department_id INT NOT NULL,
    service_name VARCHAR(255) NOT NULL,
    description TEXT,
    basic_rate DECIMAL(10, 2) NOT NULL, -- Base price of the test
    duration_minutes INT,               -- Estimated time for the test in minutes
    preparation_instructions TEXT,      -- Instructions for patients, if any
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (department_id) REFERENCES Department(department_id)
);
#user
CREATE TABLE USER (
    user_id INT PRIMARY KEY AUTO_INCREMENT,
    NAME VARCHAR(100),
    email VARCHAR(100) UNIQUE NOT NULL,
    PASSWORD VARCHAR(255),
    username VARCHAR(255) UNIQUE,
    role ENUM('admin', 'doctor', 'nurse', 'receptionist', 'patient', 'user') NOT NULL,
    phone_number VARCHAR(15),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    lab_id INT NOT NULL, 
    FOREIGN KEY (lab_id) REFERENCES Lab(lab_id)  -- removed comma and fixed reference syntax
);

CREATE TABLE Patient (
    patient_id INT PRIMARY KEY AUTO_INCREMENT,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    age INT NOT NULL,
    gender VARCHAR(10) NOT NULL,
    contact_number VARCHAR(15),
    email VARCHAR(100),
    address TEXT,
    patient_code VARCHAR(50) UNIQUE,
    user_id INT,  -- make it foriegn key afterwords . currently its normal field
    medical_history TEXT,
    blood_type VARCHAR(3),
    insurance_details TEXT,
    state_id INT,
    country_id INT,
    city_id INT,
    lab_id INT NOT NULL, 
    FOREIGN KEY (lab_id) REFERENCES lab(lab_id);
    FOREIGN KEY (state_id) REFERENCES State(state_id),
    FOREIGN KEY (country_id) REFERENCES Country(country_id),
    FOREIGN KEY (city_id) REFERENCES City(city_id)
);

# get departmens of  a branch
SELECT d.`department_name`,d.`department_id`,b.`branch_id` FROM department d 
	INNER JOIN department b ON d.branch_id = b.branch_id 
	WHERE d.branch_id = 1 ;