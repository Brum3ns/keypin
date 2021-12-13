![GitHub Light](keypin.png)
Keypin is a lightweight and highly customization tool, built to bypass forbidden pages. It supports the most common bypass techniques and also combined/adjust theses techniques for ore deep testing. 


# Features
* Friendly configuration and customaziation
* Supports diffirent techniques such as Verb, headers, path
* Adjust techniques depending on it's behavior for better detection 
* Bypass cached pages to avoid false negatives

## Versions

- v1.1 [**InProgress**] Will detect *Webhosting service*, *cache detection* and *better request preformance to handle large amount of urls in a hige preformance speed without getting rate limited or IP blocked.* 
- v1.0 Keypin works with all it's main options and does it's job as a 403/401 bypass tool.

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

**Keypin help menu. Displays all options that are available**
```
./keypin -h

```
**Simple bypass mode**

:key: Effective and most common way to use. This will run all default scans and combines techniques.
```
./keypin -u https://www.example.com -p /admin

```

**Attacking with custom Verb (HTTP method) and static header**

:key: Can be used if an early recon has been done and the user know that "X-Forward-For" is a valid supported header etc.
```
./keypin -u https://www.example.com -p /admin -H "X-Forward-For: 127.0.0.1" -m GET

```
**Attacking a forbidden website on the root without a path given**

:key: If the root page is forbidden. This scan can be used to bypass the forbidden domain when no path is known.
```
./keypin -u https://www.example.com
```

**Debugg and response information**

:key: Use for Debugging mode or to better detect response behavior from the target domain.
```
./keypin -u https://www.example.com -p /admin -v

```
