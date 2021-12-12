Keypin is a lightweight and highly customization tool, built to bypass forbidden pages. It supports the most common bypass techniques and also combined/adjust theses techniques for ore deep testing. 

# Features

* Friendly configuration and customaziation
* Supports diffirent techniques such as Verb, headers, path
* Adjust techniques depending on it's behavior for better detection 
* Bypass cached pages to avoid false negatives

# Installation

Using Golang
```
go get
```
Using Git
```
git clone https://github.com/Brum3ns/keypin.git
```

# Usage

Keypin help menu
```
./keypin -h
```
Simple bypass mode
```
./keypin -u https://www.example.com -p /admin
```
Attacking with custom Verb (HTTP method) and static header
```
./keypin -u https://www.example.com -p /admin -H "X-Forward-For: 127.0.0.1" -m GET
```
Attacking a forbidden website on the root without a path
```
./keypin -u https://www.example.com
```
Debugging mode
```
./keypin -u https://www.example.com -p /admin -v
```
