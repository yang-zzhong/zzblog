
+------------+

urlid: goproxy-intro

title: GoProxy Introduction

cate: tech

overview: The GoProxy is a high-performance http proxy, https proxy, socks5 proxy, ss proxy, websocket proxies, tcp proxies, udp proxies,  game shield, game proxies. Support forward proxies, reverse proxy, transparent proxy, internet nat proxies, https proxy load balancing, http proxy load balancing , socks5 proxies load balancing, socket proxy load balancing, ss proxy load balancing, TCP / UDP port mapping, SSH transit, TLS encrypted transmission, protocol conversion, anti-pollution DNS proxy, API authentication, speed limit, limit connection. Reverse proxies to help you expose a local server behind a NAT or firewall to the internet so that you or your visitors can access it directly and easily.

tags: #golang, #proxy

image: https://github.com/snail007/proxy_admin_free/raw/master/res/images/socks5_en.gif

lang: en

+------------+

## GOPROXY Introduction
<img src="https://github.com/snail007/goproxy/blob/master/doc/images/logo.jpg?raw=true" width="200" height="auto"/>
The GoProxy is a high-performance http proxy, https proxy, socks5 proxy, ss proxy, websocket proxies, tcp proxies, udp proxies,  game shield, game proxies. Support forward proxies, reverse proxy, transparent proxy, internet nat proxies, https proxy load balancing, http proxy load balancing , socks5 proxies load balancing, socket proxy load balancing, ss proxy load balancing, TCP / UDP port mapping, SSH transit, TLS encrypted transmission, protocol conversion, anti-pollution DNS proxy, API authentication, speed limit, limit connection. Reverse proxies to help you expose a local server behind a NAT or firewall to the internet so that you or your visitors can access it directly and easily.

---

[![stable](https://img.shields.io/badge/stable-stable-green.svg)](https://github.com/snail007/goproxy/) [![license](https://img.shields.io/github/license/snail007/goproxy.svg?style=plastic)]() [![download_count](https://img.shields.io/github/downloads/snail007/goproxy/total.svg?style=plastic)](https://github.com/snail007/goproxy/releases) [![download](https://img.shields.io/github/release/snail007/goproxy.svg?style=plastic)](https://github.com/snail007/goproxy/releases) 

---

### [点击我观看视频教程](https://space.bilibili.com/472844633)
- [中文说明](https://github.com/snail007/goproxy/blob/master/README_ZH.md)
- [中文手册](https://snail007.github.io/goproxy/manual/zh/)
- [Download](https://github.com/snail007/goproxy/releases)
- [Desktop Edition, Web Console : ProxyAdmin](https://github.com/snail007/proxy_admin_free)
- [Android Global Proxy Edition](https://github.com/snail007/goproxy-ss-plugin-android) 
- [Android Proxy Server Edition](https://github.com/snail007/goproxy-android) 
- [SDK](https://github.com/snail007/goproxy-sdk)
- [GORPOXY Manual](https://snail007.github.io/goproxy/manual/)
- [GORPOXY Tutorial](https://snail007.github.io/goproxy)
- [Free version VS commercial version](https://snail007.github.io/goproxy/free_vs_commercial/)

### ProxyAdmin Demo
And ProxyAdmin is a powerful web console of snail007/goproxy .

![](https://github.com/snail007/proxy_admin_free/raw/master/res/images/socks5_en.gif)

### What can it do?
- Chained proxies, the program itself can be used as an proxies, and if it is set up, it can be used as a secondary proxies or even an N-level proxies.
- Communication encryption, if the program is not a level one proxies, and the upper level proxies is also the program, then the communication between the upper level proxies and the upper level proxies can be encrypted, and the underlying tls high-intensity encryption is used, and the security is featureless.
- Smart HTTP, SOCKS5 proxy, will automatically determine whether the visited website is blocked. If it is blocked, it will use the upstream proxies (provided that the upstream proxies is configured) to access the website; if the visited website is not blocked, in order to speed up the access, the proxies will Direct access to the website without using a upstream proxies.
- Domain name black and white list, more free to control the way the website is accessed.
- Cross-platform, whether you are widows, linux, mac, or even raspberry pie, you can run the proxy very well.
- Multi-protocol support, support for HTTP(S), TCP, UDP, Websocket, SOCKS5 proxy.
- TCP/UDP port forwarding.
- Support intranet penetration, protocol supports TCP and UDP.
- SSH relay, HTTP (S), SOCKS5 proxy supports SSH relay, the upper Linux server does not need any server, a local proxy can be happy online.
- [KCP](https://github.com/xtaci/kcp-go) protocol support, HTTP(S), SOCKS5, SPS proxy supports KCP protocol to transmit data, reduce latency and improve browsing experience.
- Dynamic selection of upstream proxies, through the external API, HTTP (S), SOCKS5, SPS proxies can achieve user-based or IP-based speed limit, connection limit, dynamic access to upstream.
- Flexible upstream allocation, HTTP(S), SOCKS5 proxy can implement user- or IP-based speed limit, connection limit, and upper-level through configuration files.
- Transparent HTTP (S) proxy, in conjunction with iptables, forwards the outgoing 80, 443 traffic directly to the proxy at the gateway, enabling non-aware intelligent router proxy.
- Protocol conversion, which can convert existing HTTP(S) or SOCKS5 or SS proxy into one port and support HTTP(S) and SOCKS5 and SS proxy at the same time. Converted SOCKS5 and SS proxy. If the upstream is SOCKS5 proxy, then UDP is supported. Features while supporting powerful cascading authentication.
- Custom underlying encrypted transmission, http(s)\sps\socks proxy can encrypt tcp data via tls standard encryption and kcp protocol on top of tcp, in addition to support custom encryption after tls and kcp, that is Said custom encryption and tls|kcp can be used in combination, the internal AES256 encryption, you only need to define a password when you use it.
- Underlying compression efficient transmission, http(s)\sps\socks proxy can encrypt tcp data through custom encryption and tls standard encryption and kcp protocol on tcp, and can also compress data after encryption, that is, compression function And custom encryption and tls|kcp can be used in combination.
- Secure DNS proxy, which can secure and prevent pollution DNS queries through encrypted proxy communication between the DNS proxy server provided by the local proxy and the upstream proxy.
- Load balancing, high availability, HTTP(S)\SOCKS5\SPS proxies supports upstream load balancing and high availability, and multiple upstream repeat-P parameters can be used.
- Specify the egress IP. The HTTP(S)\SOCKS5\SPS\TCP proxy supports the client to connect with the ingress IP, and uses the ingress IP as the egress IP to access the target website. If the ingress IP is an intranet IP, the egress IP does not use the ingress IP.
- Support speed limit, HTTP(S)\SOCKS5\SPS\TCP proxy supports speed limit.
- SOCKS5 proxies supports cascading certification.
- The certificate parameter uses base64 data. By default, the -C, -K parameter is the path of the crt certificate and the key file. If it is the beginning of base64://, then the latter data is considered to be base64 encoded and will be used after decoding.
- Support client IP black and white list, more secure control of client access to proxy service, if black and white list is set at the same time, then only whitelist is effective. Socks / HTTP(S) / SPS / TCP / UDP / DNS / intranet NAT The bridge/intranet NAT the tbridge and supports the client IP black and white list.
- Range ports listen on, HTTP(S)\SOCKS5\SPS\TCP proxy supports port range listening, avoiding starting too many processes and improving performance.

### Why do you need it?

- When for some reason we are unable to access our services elsewhere, we can establish a secure tunnel to access our services through multiple connected proxy nodes.
- WeChat interface is developed locally for easy debugging.
- Remote access to intranet machines.
- Play LAN games with your friends.
- I used to play only on the LAN, and now I can play anywhere.
- Replace the sword inside Netnet, show IP internal Netcom, peanut shell and other tools.
- ..

 
The manual on this page applies to the latest version of goproxy. Other versions may not be applicable. Please use the command according to your own instructions.
 

### Joining the organization
[Click to join the gitter](https://gitter.im/go-proxy/Lobby?utm_source=share-link&utm_medium=link&utm_campaign=share-link)

[Click to join the TG](https://t.me/snail007_goproxy)

## Download and install 

### Quick installation

0. If your VPS is a Linux 64-bit system, you only need to execute the following sentence to complete the automatic installation and configuration.

Tip: All operations require root privileges.

The free version performs this:

```shell
curl -L https://raw.githubusercontent.com/snail007/goproxy/master/install_auto.sh | bash
```

The commercial version performs this:

```shell
curl -L https://raw.githubusercontent.com/snail007/goproxy/master/install_auto_commercial.sh | bash
```

The installation is complete, the configuration directory is /etc/proxy. For more detailed usage, please refer to the manual directory above to learn more about the features you want to use.
If the installation fails or your vps is not a linux64-bit system, follow the semi-automatic steps below to install:
  
### Manual installation

1. Download the proxy

Download address: https://github.com/snail007/goproxy/releases/latest

Let's take v7.9 as an example. If you have the latest version, please use the latest version of the link. Note that the version number in the download link below is the latest version number.

The free version performs this:

```shell
cd /root/proxy/
w ge t https://github.com/snail007/goproxy/releases/download/v7.9/proxy-linux-amd64.tar.gz
```

The commercial version performs this:

```shell
cd /root/proxy/
wget https://github.com/snail007/goproxy/releases/download/v7.9/proxy-linux-amd64_commercial.tar.gz
```

2. Download the automatic installation script

The free version performs this:

```shell
c d /root/proxy/
wget https://raw.githubusercontent.com/snail007/goproxy/master/install.sh
Chmod +x install.sh
./install.sh
```

The commercial version performs this:

```shell
cd /root/proxy/
wget https://raw.githubusercontent.com/snail007/goproxy/master/install_commercial.sh
Chmod +x install_commercial.sh
./install_commercial.sh
```

## TODO
- http,socks proxy multiple upstream load balancing?
- http(s) proxy to increase pac support?
- Welcome to add group feedback..

## License
Proxy is licensed under GPLv3 license.

## Contact
Official QQ exchange group: 793015219

## Donation
If the proxy helps you solve a lot of problems, you can better support the proxy through the donation below.

BTC  ADDRESS: `1BJcBhGhREiz1q3VTYoiVPuAZy5PGxRG9z`

<img src="https://raw.githubusercontent.com/snail007/goproxy/master/docs/img/btc.png" width="150" height="auto"/>  

ETH  ADDRESS: `0x0fA4c567768d2E59E6221152EA52F4842866AFC8`

<img src="https://raw.githubusercontent.com/snail007/goproxy/master/docs/img/eth.png" width="150" height="auto"/>  

### Source code declaration

The author of this project found that a large number of developers based on the project for secondary development or using a large number of core code of the project without complying with the GPLv3 agreement, which seriously violates the original intention of using the GPLv3 open source agreement in this project. In view of this situation, the project adopts the source. The code delays the release strategy, to a certain extent, to curb these behaviors that do not respect open source and do not respect the labor results of others.
This project will continue to update the iterations and continue to release the full platform binary program, providing you with powerful and convenient proxies tools.
If you have customized, business needs, please send an email to `arraykeys@gmail.com`

## Goproxy Manual


## First Start  
    
### 1. Environment  

The manual tutorial, the default system is linux, the program is proxy; all operations require root privileges;  

If you are windows, please use the windows version of proxy.exe.  
    
### 2. Using configuration files  

The next tutorial will introduce the usage method through the command line parameters, or you can get the parameters by reading the configuration file.  

The specific format is to specify the configuration file by the @ symbol, for example: ./proxy @configfile.txt  

The format in configfile.txt is that the first line is the name of the subcommand, and the second line starts with one parameter per line.  

Format: `parameter Parameter value`, direct write parameter without parameter value, for example: --nolog  

For example, the contents of configfile.txt are as follows:  

```shell  
Http  
-t tcp  
-p :33080  
--forever  
```  

### 3. Debug output  

By default, the information output by the log does not include the number of file lines. In some cases, in order to troubleshoot the program, the problem is quickly located.  

You can use the --debug parameter to output the number of lines of code and milliseconds.  

### 4. Using log files  

By default, the log is displayed directly in the console. If you want to save to a file, you can use the --log parameter.  

For example: --log proxy.log, the log will be output to the proxy.log to facilitate troubleshooting.  


### 5. Generate the certificate file required for encrypted communication  

The http, tcp, udp proxy process communicates with the upstream. For security, we use encrypted communication. Of course, we can choose not to encrypt the communication. All the communication and the upstream communication in this tutorial are encrypted, and the certificate file is required.  

1. Generate a self-signed certificate and key file with the following command.  
`./proxy keygen -C proxy`  
The certificate file proxy.crt and the key file proxy.key will be generated under the current program directory.  

2. Use the following command to generate a new certificate using the self-signed certificate proxy.crt and the key file proxy.key: goproxy.crt and goproxy.key.  
`./proxy keygen -s -C proxy -c goproxy`  
The certificate file goproxy.crt and the key file goproxy.key will be generated under the current program directory.  

3. By default, the domain name inside the certificate is random and can be specified using the `-n test.com` parameter.  

4. More usage: `proxy keygen --help`.  

### 6. Running in the background  

After the proxy is executed by default, you cannot close the command line if you want to keep the proxy running.  

If you want to run the proxy in the background, the command line can be closed, just add the --daemon parameter at the end of the command.  

For example:  

`./proxy http -t tcp -p "0.0.0.0:38080" --daemon`  

### 7. Guardian running  
The daemon runs the parameter --forever, for example: `proxy http --forever` ,  

The proxy will fork the child process, and then monitor the child process. If the child process exits abnormally, restart the child process after 5 seconds.  

This parameter is matched with the background running parameter --daemon and log parameter --log, which can guarantee that the proxy will always execute in the background without accidentally exiting.  

And you can see the output log content of the proxy through the log file.  

For example: `proxy http -p ":9090" --forever --log proxy.log --daemon`  

### 8. Security advice  

When the VPS is behind the nat device, the vps network card IP is the intranet IP. At this time, you can use the -g parameter to add the vps external network ip to prevent the infinite loop.  

Suppose your vps external network ip is 23.23.23.23. The following command sets 23.23.23.23 with the -g parameter.  

`./proxy http -g "23.23.23.23"`  

### 9. Load balancing and high availability  

The HTTP(S)\SOCKS5\SPS proxy supports upper-level load balancing and high availability, and multiple upstream repeat-P parameters can be used.  

The load balancing policy supports five types, which can be specified by the `--lb-method` parameter:  

Roundrobin used in turn  

Leastconn uses the minimum number of connections  

Leasttime uses the least connection time  

Hash uses a fixed upstream based on the client address  

Weight Select a upstream according to the weight and number of connections of each upstream  

prompt:  

The load balancing check interval can be set by `--lb-retrytime` in milliseconds.  

The load balancing connection timeout can be set by `--lb-timeout` in milliseconds.  

If the load balancing policy is weight, the -P format is: 2.2.2.2: 3880?w=1, where 1 is the weight and an integer greater than 0.  

If the load balancing policy is hash, the default is to select the upstream based on the client address. You can select the upstream by using the destination address of the access `--lb-hashtarget`.  


### 10. Agent springboard jump  

Http (s) agent, SPS agent, intranet penetration, tcp agent support the connection of upstreams through intermediate third-party agents,  

The parameters are: --jumper, all the formats are as follows:  

```text  
 Http://username:password@host:port  
 Http://host:port  
 Https://username:password@host:port  
 Https://host:port  
 Socks5://username:password@host:port  
 Socks5://host:port  
 Socks5s://username:password@host:port  
 Socks5s://host:port  
 Ss://method:password@host:port  
```  

Http,socks5 represents the normal http and socks5 proxy.  

Https,socks5s represents the http and socks5 agents protected by tls.  

That is http proxy over TLS, socks over TLS.  

### 11. Domain Name Black and White List  

The socks/http(s)/sps proxy supports domain name black and white lists.  

Use the --stop parameter to specify a domain name blacklist file, then the connection will be disconnected when the user connects these domains in the file.  

Specify a domain name whitelist file with the --only parameter, then the connection will be disconnected when the user connects to a domain other than those domains in the file.  

If both --stop and --only are set, then only --only will work.  

The format of the black and white domain name list file is as follows:  

```text  
**.baidu.com  
*.taobao.com  
A.com  
192.168.1.1  
192.168.*.*  
?.qq.com  
```  

Description:  

1. One domain name per line, domain name writing supports wildcards `*` and `?`, `*` represents any number of characters, `?` represents an arbitrary character,  

2.`**.baidu.com` Matches no matter how many levels all suffixes are ..baidu.com`.  

3.`*.taobao.com` The matching suffix is ​​the third-level domain name of `.taobao.com`.  

4. It can also be an IP address directly.  

5.`#` at the beginning of the comment.  

### 12. Client IP Blacklist and Whitelist  

socks/http(s)/sps/tcp/udp/dns/ intranet penetration bridge/intranet penetration tbridge, support client IP black and white list.  

Use the --ip-deny parameter to specify a client IP blacklist list file, then the connection will be disconnected when the user's IP is in this file.  
    
Use the --ip-allow parameter to specify a client IP whitelist file, then the connection will be disconnected when the user's IP is not in the file.  

If both --ip-deny and --ip-allow are set, then only --ip-allow will work.  

The format of the client IP blacklist and whitelist file is as follows:  

```text  
192.168.1.1  
192.168.*.*  
192.168.1?.*  
```  

Description:  

1. One domain name per domain, domain name writing supports wildcards `*` and `?`, `*` represents any number of characters, `?` represents an arbitrary character.  

2.`#` at the beginning of the comment.  

### 13. Protocol loading file  

There are many places in the proxy's various proxy functions to set a file. For example: --blocked Specifies a domain name list file that goes directly to the upper level. The parameter value is the path of the file.  

If the parameter supports the protocol loading file, the file path can be not only the file path, but also:  

a. The base64 encoding at the beginning of "base64://" indicates the contents of the above file, for example: base64://ajfpoajsdfa=  

b. "str://" at the beginning of the English comma separated multiple, such as: str://xxx, yyy  

The proxy's blocked, direct, stop, only, hosts, resolve.rules, rewriter.rules, ip.allow, ip.deny files support protocol loading.  

## 1.HTTP Proxies  

### 1.1. Ordinary level HTTP proxy  

![1.1](https://raw.githubusercontent.com/snail007/goproxy/master/doc/images/http-1.png)  

`./proxy http -t tcp -p "0.0.0.0:38080"`  
    
Listen port argument `-p` can be:

```text
  -p ":8081"  listen on 8081
  -p ":8081,:8082"  listen on 8081 and 8082
  -p ":8081,:8082,:9000-9999" listen on 8081 and 8082 and 9000 and 9001 to 9999, 1002 total ports  
```

### 1.2. Ordinary secondary HTTP proxy  

![1.2](https://raw.githubusercontent.com/snail007/goproxy/master/doc/images/http-2.png)  

Use local port 8090, assuming the upstream HTTP proxy is `22.22.22.22:8080`  

`./proxy http -t tcp -p "0.0.0.0:8090" -T tcp -P "22.22.22.22:8080" `  

We can also specify the black and white list file of the website domain name, one domain name per line, the matching rule is the rightmost match, for example: baidu.com, the match is *.*.baidu.com, the blacklist domain name goes directly to the upstream agent, whitelist The domain name does not go to the upstream agent.  

`./proxy http -p "0.0.0.0:8090" -T tcp -P "22.22.22.22:8080" -b blocked.txt -d direct.txt`  
    
### 1.3.HTTP secondary agent (encryption)  

> Note: The `proxy.crt` and `proxy.key` used by the secondary proxy should be consistent with the primary proxy.  

![1.3](https://raw.githubusercontent.com/snail007/goproxy/master/doc/images/http-tls-2.png)  
Level 1 HTTP proxy (VPS, IP: 22.22.22.22)  
`./proxy http -t tls -p ":38080" -C proxy.crt -K proxy.key`  
    
Secondary HTTP proxy (local Linux)  
`./proxy http -t tcp -p ":8080" -T tls -P "22.22.22.22:38080" -C proxy.crt -K proxy.key`  
Then access the local port 8080 is to access the proxy port 38080 on the VPS.  
    
Secondary HTTP proxy (local windows)  
`./proxy.exe http -t tcp -p ":8080" -T tls -P "22.22.22.22:38080" -C proxy.crt -K proxy.key`  
Then set your windos system, the proxy that needs to go through the proxy Internet program is http mode, the address is: 127.0.0.1, the port is: 8080, the program can access the Internet through vps through the encrypted channel.  
    
### 1.4.HTTP Level 3 Agent (Encryption)  
![1.3](https://raw.githubusercontent.com/snail007/goproxy/master/doc/images/http-tls-3.png)  
Level 1 HTTP proxy VPS_01, IP: 22.22.22.22  
`./proxy http -t tls -p ":38080" -C proxy.crt -K proxy.key`  
Secondary HTTP proxy VPS_02, IP: 33.33.33.33  
`./proxy http -t tls -p ":28080" -T tls -P "22.22.22.22:38080" -C proxy.crt -K proxy.key`  
Level 3 HTTP proxy (local)  
`./proxy http -t tcp -p ":8080" -T tls -P "33.33.33.33:28080" -C proxy.crt -K proxy.key`  
Then accessing the local port 8080 is to access the proxy port 38080 on the primary HTTP proxy.  
    
### 1.5.Basic certification  
For the proxy HTTP protocol, we can perform Basic authentication. The authenticated username and password can be specified on the command line.  
`./proxy http -t tcp -p ":33080" -a "user1:pass1" -a "user2:pass2"`  
For multiple users, repeat the -a parameter.  
It can also be placed in a file in the format of a "username:password" and then specified with -F.  
`./proxy http -t tcp -p ":33080" -F auth-file.txt`  
    
In addition, the http(s) proxy also integrates external HTTP API authentication. We can specify an http url interface address with the --auth-url parameter.  
Then when there is a user connection, the proxy will request the url in GET mode, and bring the following four parameters. If the HTTP status code 204 is returned, the authentication is successful.  
In other cases, the authentication failed.  
For example:  
`./proxy http -t tcp -p ":33080" --auth-url "http://test.com/auth.php"`  
When the user connects, the proxy will request the url ("http://test.com/auth.php") in GET mode.  
Take five parameters: user, pass, ip, local_ip, target:  
Http://test.com/auth.php?user={USER}&pass={PASS}&ip={IP}&local_ip={LOCAL_IP}&target={TARGET}  
User: username  
Pass: password  
Ip: User's IP, for example: 192.168.1.200  
Local_ip: IP of the server accessed by the user, for example: 3.3.3.3  
Target: URL accessed by the user, for example: http://demo.com:80/1.html or https://www.baidu.com:80  

If there is no -a or -F or --auth-url parameter, the Basic authentication is turned off.  

### 1.6. HTTP proxy traffic is forced to go to the upper HTTP proxy  
By default, the proxy will intelligently determine whether a website domain name is inaccessible. If it is not accessible, it will go to the upper level HTTP proxy. With --always, all HTTP proxy traffic can be forced to go to the upper HTTP proxy.  
`./proxy http --always -t tls -p ":28080" -T tls -P "22.22.22.22:38080" -C proxy.crt -K proxy.key`  

### 1.7.HTTP(S) via SSH relay  
![1.7](https://raw.githubusercontent.com/snail007/goproxy/master/doc/images/http-ssh-1.png)  
Description: The principle of ssh transfer is to use the forwarding function of ssh, that is, after you connect to ssh, you can access the target address through ssh proxy.  
Suppose there is: vps  
- IP is 2.2.2.2, ssh port is 22, ssh username is: user, ssh user password is: demo  
- The user's ssh private key name is user.key  

#### *1.7.1 How to ssh username and password*  
Local HTTP(S) proxy port 28080, executing:  
`./proxy http -T ssh -P "2.2.2.2:22" -u user -D demo -t tcp -p ":28080"`  
#### *1.7.2 How to ssh username and key*  
Local HTTP(S) proxy port 28080, executing:  
`./proxy http -T ssh -P "2.2.2.2:22" -u user -S user.key -t tcp -p ":28080"`  

### 1.8.KCP protocol transmission  
![1.8](https://raw.githubusercontent.com/snail007/goproxy/master/doc/images/http-kcp.png)  
The KCP protocol requires the --kcp-key parameter to set a password for encrypting and decrypting data.  

Level 1 HTTP proxy (VPS, IP: 22.22.22.22)  
`./proxy http -t kcp -p ":38080" --kcp-key mypassword`  
    
Secondary HTTP proxy (local Linux)  
`./proxy http -t tcp -p ":8080" -T kcp -P "22.22.22.22:38080" --kcp-key mypassword`  
Then access the local port 8080 is to access the proxy port 38080 on the VPS, the data is transmitted through the kcp protocol, note that the kcp is the udp protocol, so the firewall needs to release the 380p udp protocol.  

### 1.9 HTTP(S) Reverse Proxy  
![1.9](https://raw.githubusercontent.com/snail007/goproxy/master/doc/images/fxdl.png)  
The proxy not only supports the proxy setting in other software, but also provides proxy services for other software. It also supports directly parsing the requested website domain name to the proxy listening ip, and then the proxy listens to the 80 and 443 ports, then the proxy will automatically You proxy access to the HTTP(S) website you need to access.  

How to use:  
On the "last level proxy proxy" machine, because the proxy is to be disguised as all websites, the default HTTP port of the website is 80, HTTPS is 443, and the proxy can listen to ports 80 and 443. Parameters -p multiple addresses with commas segmentation.  
`./proxy http -t tcp -p :80,:443`  

This command starts a proxy agent on the machine, and listens to ports 80 and 443 at the same time. It can be used as a normal proxy, or directly resolve the domain name that needs to be proxyed to the IP of this machine.  

If there is a upstream agent, then refer to the above tutorial to set the upstream, the use is exactly the same.  
`./proxy http -t tcp -p :80,:443 -T tls -P "2.2.2.2:33080" -C proxy.crt -K proxy.key`  

Note:  
The DNS resolution result of the server where the proxy is located cannot be affected by the custom resolution, otherwise it will be infinite loop. The proxy proxy should specify the `--dns 8.8.8.8` parameter.  

### 1.10 HTTP(S) Transparent Proxy  
This mode needs to have a certain network foundation. If the related concepts are not understood, please search for it yourself.  
Assuming the proxy is now running on the router, the startup command is as follows:  
`./proxy http -t tcp -p :33080 -T tls -P "2.2.2.2:33090" -C proxy.crt -K proxy.key`  

Then add the iptables rule, here are the reference rules:  
```shell  
#Upper proxy server IP address:  
Proxy_server_ip=2.2.2.2  

#路由器Running port for proxy listening:  
Proxy_local_port=33080  

#The following does not need to be modified  
#create a new chain named PROXY  
Iptables -t nat -N PROXY  

# Ignore your PROXY server's addresses  
# It's very IMPORTANT, just be careful.  

Iptables -t nat -A PROXY -d $proxy_server_ip -j RETURN  

# Ignore LANs IP address  
Iptables -t nat -A PROXY -d 0.0.0.0/8 -j RETURN  
Iptables -t nat -A PROXY -d 10.0.0.0/8 -j RETURN  
Iptables -t nat -A PROXY -d 127.0.0.0/8 -j RETURN  
Iptables -t nat -A PROXY -d 169.254.0.0/16 -j RETURN  
Iptables -t nat -A PROXY -d 172.16.0.0/12 -j RETURN  
Iptables -t nat -A PROXY -d 192.168.0.0/16 -j RETURN  
Iptables -t nat -A PROXY -d 224.0.0.0/4 -j RETURN  
Iptables -t nat -A PROXY -d 240.0.0.0/4 -j RETURN  

# Anything to port 80 443 should be redirected to PROXY's local port  
Iptables -t nat -A PROXY -p tcp --dport 80 -j REDIRECT --to-ports $proxy_local_port  
Iptables -t nat -A PROXY -p tcp --dport 443 -j REDIRECT --to-ports $proxy_local_port  

# Apply the rules to nat client  
Iptables -t nat -A PREROUTING -p tcp -j PROXY  
# Apply the rules to localhost  
Iptables -t nat -A OUTPUT -p tcp -j PROXY  
```  
- Clear the entire chain iptables -F Chain names such as iptables -t nat -F PROXY  
- Delete the specified user-defined chain iptables -X chain name such as iptables -t nat -X PROXY  
- Remove rules from the selected chain iptables -D chain name Rule details such as iptables -t nat -D PROXY -d 223.223.192.0/255.255.240.0 -j RETURN  

### 1.11 Custom DNS  
--dns-address and --dns-ttl parameters, used to specify the dns (--dns-address) used by the proxy to access the domain name.  
And the analysis result cache time (--dns-ttl) seconds, to avoid system dns interference to the proxy, in addition to the cache function can also reduce the dns resolution time to improve access speed.  
For example:  
`./proxy http -p ":33080" --dns-address "8.8.8.8:53" --dns-ttl 300`  

### 1.12 Custom encryption  
The proxy's http(s) proxy can encrypt tcp data via tls standard encryption and kcp protocol on top of tcp, in addition to support customization after tls and kcp.  
Encryption, that is to say, custom encryption and tls|kcp can be used in combination. The internal use of AES256 encryption, you only need to define a password when you use it.  
Encryption is divided into two parts, one is whether the local (-z) encryption and decryption, and the other is whether the transmission with the upstream (-Z) is encrypted or decrypted.  
Custom encryption requires both ends to be proxy. The following two levels and three levels are used as examples:  

Secondary instance  

Execute on level 1 vps (ip: 2.2.2.2):  
`proxy http -t tcp -z demo_password -p :7777`  
Local secondary execution:  
`proxy http -T tcp -P 2.2.2.2:777 -Z demo_password -t tcp -p :8080`  
In this way, when the website is accessed through the local agent 8080, the target website is accessed through encrypted transmission with the upstream.  


Three-level instance  

Execute on level 1 vps (ip: 2.2.2.2):  
`proxy http -t tcp -z demo_password -p :7777`  
Execute on the secondary vps (ip: 3.3.3.3):  
`proxy http -T tcp -P 2.2.2.2:7777 -Z demo_password -t tcp -z other_password -p :8888`  
Local three-level execution:  
`proxy http -T tcp -P 3.3.3.3:8888 -Z other_password -t tcp -p :8080`  
In this way, when the website is accessed through the local agent 8080, the target website is accessed through encrypted transmission with the upstream.  

### 1.13 Compressed transmission  
The proxy http(s) proxy can encrypt tcp data through tls standard encryption and kcp protocol on top of tcp, and can also compress data before custom encryption.  
That is to say, compression and custom encryption and tls|kcp can be used in combination. Compression is divided into two parts, one part is local (-m) compression transmission.  
Part of it is compressed with the upstream (-M) transmission.  
Compression requires both sides to be proxy. Compression also protects (encrypted) data to a certain extent. The following uses Level 2 and Level 3 as examples:  

Secondary instance  

Execute on level 1 vps (ip: 2.2.2.2):  
`proxy http -t tcp -m -p :7777`  
Local secondary execution:  
`proxy http -T tcp -P 2.2.2.2:777 -M -t tcp -p :8080`  
In this way, when the website is accessed through the local agent 8080, the target website is accessed through compression with the upstream.  


Three-level instance  

Execute on level 1 vps (ip: 2.2.2.2):  
`proxy http -t tcp -m -p :7777`  
Execute on the secondary vps (ip: 3.3.3.3):  
`proxy http -T tcp -P 2.2.2.2:7777 -M -t tcp -m -p :8888`  
Local three-level execution:  
`proxy http -T tcp -P 3.3.3.3:8888 -M -t tcp -p :8080`  
In this way, when the website is accessed through the local agent 8080, the target website is accessed through compression with the upstream.  

### 1.14 Load Balancing  

The HTTP(S) proxy supports upper-level load balancing, and multiple upstream repeat-P parameters can be used.  

`proxy http --lb-method=hash -T tcp -P 1.1.1.1:33080 -P 2.1.1.1:33080 -P 3.1.1.1:33080`  

### 1.14.1 Setting the retry interval and timeout time  

`proxy http --lb-method=leastconn --lb-retrytime 300 --lb-timeout 300 -T tcp -P 1.1.1.1:33080 -P 2.1.1.1:33080 -P 3.1.1.1:33080 -t tcp - p :33080`  

### 1.14.2 Setting weights  

`proxy http --lb-method=weight -T tcp -P 1.1.1.1:33080?w=1 -P 2.1.1.1:33080?w=2 -P 3.1.1.1:33080?w=1 -t tcp - p :33080`  

### 1.14.3 Use the target address to select the upstream  

`proxy http --lb-hashtarget --lb-method=hash -T tcp -P 1.1.1.1:33080 -P 2.1.1.1:33080 -P 3.1.1.1:33080 -t tcp -p :33080`  

### 1.15 Speed ​​limit  

The speed limit is 100K, which can be specified by the `-l` parameter, for example: 100K 2000K 1M . 0 means no limit.  

`proxy http -t tcp -p 2.2.2.2:33080 -l 100K`  

### 1.16 Specifying Outgoing IP  

The `--bind-listen` parameter can be used to open the client connection with the portal IP, and use the portal IP as the outgoing IP to access the target website. If the incorrect IP is bound, the proxy will not work. At this point, the proxy will try to bind the target without binding the IP, and the log will prompt.  

`proxy http -t tcp -p 2.2.2.2:33080 --bind-listen`  

### 1.17 Certificate parameters use base64 data  

By default, the -C, -K parameter is the path to the crt certificate and the key file.  

If it is the beginning of base64://, then the latter data is considered to be base64 encoded and will be used after decoding.  

### 1.18 Intelligent mode  
Intelligent mode setting, can be one of intelligent|direct|parent.  
The default is: intelligent.  
The meaning of each value is as follows:  
`--intelligent=direct`, the targets in the blocked are not directly connected.  
`--intelligent=parent`, the target that is not in the direct is going to the higher level.  
`--intelligent=intelligent`, blocked and direct have no targets, intelligently determine whether to use the upstream access target.  

### 1.19 View help  
`./proxy help http`  

## 2.TCP Proxies  
    
### 2.1. Ordinary level TCP proxy  
![2.1](https://raw.githubusercontent.com/snail007/goproxy/master/doc/images/tcp-1.png)  
Local execution:  
`./proxy tcp -p ":33080" -T tcp -P "192.168.22.33:22"`  
Then access the local port 33080 is to access port 22 of 192.168.22.33.  
    
The `-p` parameter supports :

```text
  -p ":8081" listen on 8081
  -p ":8081,:8082" listen on 8081 and 8082
  -p ":8081,:8082,:9000-9999" listen on 8081 and 8082 and 9000, 9001 to 9999 for a total of 1002 ports
```

If the number of local listening ports is greater than 1, the corresponding upper port corresponding to the local port will be connected, and the port in `-P` will be ignored.

If you need a connection from all ports, connect to the upper specified port, you can add the parameter `--lock-port`.

such as:

`./proxy tcp -p ":33080-33085" -T tcp -P "192.168.22.33:0"`

Then the connection of the `33080` port will connect to the `33080` port of 192.168.22.33, and the other ports are similar. The local and upper ports are the same. At this time, the port in the parameter `-P` uses `0`.

If you want to connect the ports of `33080`, `33081`, etc. to the `22` port of 192.168.22.33, you can add the parameter `--lock-port`.

`./proxy tcp -p ":33080-33085" -T tcp -P "192.168.22.33:22" --lock-port`

### 2.2. Ordinary secondary TCP proxy  
![2.2](https://raw.githubusercontent.com/snail007/goproxy/master/doc/images/tcp-2.png)  
VPS (IP: 22.22.2.33) is executed:  
`./proxy tcp -p ":33080" -T tcp -P "127.0.0.1:8080"`  
Local execution:  
`./proxy tcp -p ":23080" -T tcp -P "22.22.22.33:33080"`  
Then access the local port 23080 is to access port 8020 of 22.22.22.33.  
    
### 2.3. Ordinary three-level TCP proxy  
![2.3](https://raw.githubusercontent.com/snail007/goproxy/master/doc/images/tcp-3.png)  
Primary TCP proxy VPS_01, IP: 22.22.22.22  
`./proxy tcp -p ":38080" -T tcp -P "66.66.66.66:8080"`  
Secondary TCP proxy VPS_02, IP: 33.33.33.33  
`./proxy tcp -p ":28080" -T tcp -P "22.22.22.22:38080"`  
Level 3 TCP proxy (local)  
`./proxy tcp -p ":8080" -T tcp -P "33.33.33.33:28080"`  
Then access the local port 8080 is to access the port 8080 of 66.66.66.66 through the encrypted TCP tunnel.  
    
### 2.4. Encrypting secondary TCP proxy  
![2.4](https://raw.githubusercontent.com/snail007/goproxy/master/doc/images/tcp-tls-2.png)  
VPS (IP: 22.22.2.33) is executed:  
`./proxy tcp -t tls -p ":33080" -T tcp -P "127.0.0.1:8080" -C proxy.crt -K proxy.key`  
Local execution:  
`./proxy tcp -p ":23080" -T tls -P "22.22.22.33:33080" -C proxy.crt -K proxy.key`  
Then access the local port 23080 is to access the port 8080 of 22.22.22.33 through the encrypted TCP tunnel.  
    
### 2.5.Encrypting Level 3 TCP Agent  
![2.5](https://raw.githubusercontent.com/snail007/goproxy/master/doc/images/tcp-tls-3.png)  
Primary TCP proxy VPS_01, IP: 22.22.22.22  
`./proxy tcp -t tls -p ":38080" -T tcp -P "66.66.66.66:8080" -C proxy.crt -K proxy.key`  
Secondary TCP proxy VPS_02, IP: 33.33.33.33  
`./proxy tcp -t tls -p ":28080" -T tls -P "22.22.22.22:38080" -C proxy.crt -K proxy.key`  
Level 3 TCP proxy (local)  
`./proxy tcp -p ":8080" -T tls -P "33.33.33.33:28080" -C proxy.crt -K proxy.key`  
Then access the local port 8080 is to access the port 8080 of 66.66.66.66 through the encrypted TCP tunnel.  
    
### 2.6 Connecting to a upstream through a proxy  
Sometimes the network where the proxy is located cannot directly access the external network. You need to use an https or socks5 proxy to access the Internet. Then this time  
The -J parameter can help you to connect the proxy to the peer-P through the https or socks5 proxy when mapping the proxy tcp port, mapping the external port to the local.  
The -J parameter format is as follows:  

Https proxy writing:  
The proxy needs authentication, username: username password: password  
Https://username:password@host:port  
Agent does not require authentication  
Https://host:port  

Socks5 proxy writing:  
The proxy needs authentication, username: username password: password  
Socks5://username:password@host:port  
Agent does not require authentication  
Socks5://host:port  

Host: the IP or domain name of the proxy  
Port: the port of the proxy  

### 2.7 Specify Outgoing IP
When the TCP proxy is a superior type (parameter: -T) is tcp, it supports the specified exit IP. Using the `--bind-listen` parameter, you can open the client to connect with the portal IP, and use the portal IP as the outgoing IP to access the target website. If an incorrect IP is bound, the proxy will not work, the proxy will try to bind the target without binding the IP, and the log will prompt.

`./proxy tcp -p ":33080" -T tcp -P" 192.168.22.33:22" -B`

### 2.8 Speed ​​limit, connections limit

The parameter `--max-conns` can limit the maximum number of connections per port.  
For example, limit the maximum number of connections per port:  
`./proxy tcp -p ":33080" -T tcp -P "192.168.22.33:22" --max-conns 1000`  
The parameter `--rate-limit` can limit the rate of each tcp connection.  
For example, limit the connection rate of each tcp to 100k/s:  
`./proxy tcp -p ":33080" -T tcp -P "192.168.22.33:22" --rate-limit 100k`  

### 2.9 View Help
`./proxy help tcp`  

## 3.UDP Proxies  
    
### 3.1. Ordinary UDP proxy  
![3.1](https://raw.githubusercontent.com/snail007/goproxy/master/doc/images/udp-1.png)  
Local execution:  
`./proxy udp -p ":5353" -T udp -P "8.8.8.8:53"`  
Then access the local UDP: 5353 port is to access 8.8.8.8 UDP: 53 port.  
  
The `-p` parameter supports :

```text
  -p ":8081" listen on 8081
  -p ":8081,:8082" listen on 8081 and 8082
  -p ":8081,:8082,:9000-9999" listen on 8081 and 8082 and 9000, 9001 to 9999 for a total of 1002 ports
```

If the number of local listening ports is greater than 1, the corresponding upper port corresponding to the local port will be connected, and the port in `-P` will be ignored.

If you need a connection from all ports, connect to the upper specified port, you can add the parameter `--lock-port`.

such as:

`./proxy udp -p ":33080-33085" -T udp -P "192.168.22.33:0"`

Then the connection of the `33080` port will connect to the `33080` port of 192.168.22.33, and the other ports are similar. The local and upper ports are the same. At this time, the port in the parameter `-P` uses `0`.

If you want to connect the ports of `33080`, `33081`, etc. to the `2222` port of 192.168.22.33, you can add the parameter `--lock-port`.

`./proxy udp -p ":33080-33085" -T udp -P "192.168.22.33:2222" --lock-port`

### 3.2. Ordinary secondary UDP proxy  
![3.2](https://raw.githubusercontent.com/snail007/goproxy/master/doc/images/udp-2.png)  
VPS (IP: 22.22.2.33) is executed:  
`./proxy tcp -p ":33080" -T udp -P "8.8.8.8:53"`  
Local execution:  
`./proxy udp -p ":5353" -T tcp -P "22.22.22.33:33080"`  
Then access the local UDP: 5353 port is through the TCP tunnel, through the VPS access 8.8.8.8 UDP: 53 port.  
    
### 3.3. Ordinary three-level UDP proxy  
![3.3](https://raw.githubusercontent.com/snail007/goproxy/master/doc/images/udp-3.png)  
Primary TCP proxy VPS_01, IP: 22.22.22.22  
`./proxy tcp -p ":38080" -T udp -P "8.8.8.8:53"`  
Secondary TCP proxy VPS_02, IP: 33.33.33.33  
`./proxy tcp -p ":28080" -T tcp -P "22.22.22.22:38080"`  
Level 3 TCP proxy (local)  
`./proxy udp -p ":5353" -T tcp -P "33.33.33.33:28080"`  
Then access to the local 5353 port is through the TCP tunnel, through the VPS to access port 8.8.8.8.  
    
### 3.4. Encrypting secondary UDP proxy  
![3.4](https://raw.githubusercontent.com/snail007/goproxy/master/doc/images/udp-tls-2.png)  
VPS (IP: 22.22.2.33) is executed:  
`./proxy tcp -t tls -p ":33080" -T udp -P "8.8.8.8:53" -C proxy.crt -K proxy.key`  
Local execution:  
`./proxy udp -p ":5353" -T tls -P "22.22.22.33:33080" -C proxy.crt -K proxy.key`  
Then access the local UDP: 5353 port is through the encrypted TCP tunnel, through the VPS access 8.8.8.8 UDP: 53 port.  
    
### 3.5. Encryption Level 3 UDP Agent  
![3.5](https://raw.githubusercontent.com/snail007/goproxy/master/doc/images/udp-tls-3.png)  
Primary TCP proxy VPS_01, IP: 22.22.22.22  
`./proxy tcp -t tls -p ":38080" -T udp -P "8.8.8.8:53" -C proxy.crt -K proxy.key`  
Secondary TCP proxy VPS_02, IP: 33.33.33.33  
`./proxy tcp -t tls -p ":28080" -T tls -P "22.22.22.22:38080" -C proxy.crt -K proxy.key`  
Level 3 TCP proxy (local)  
`./proxy udp -p ":5353" -T tls -P "33.33.33.33:28080" -C proxy.crt -K proxy.key`  
Then access the local 5353 port is to access the 8.8.8.8 port 53 through VPS_01 through the encrypted TCP tunnel.  
    
### 3.6. View help  
`./proxy help udp`  

## 4. Internet NAT  
### 4.1 principle description  
Intranet penetration, divided into two versions, "multi-link version" and "multiplexed version", generally like a web service, this service is not a long-term connection, it is recommended to use "multi-link version", if it is to keep long The time connection suggests using a "multiplexed version."  
1. Multi-link version, the corresponding sub-command is tserver, tclient, tbridge.  
1. Multiplexed version, the corresponding subcommand is server, client, bridge.  
1. The parameters of the multi-link version and the multiplex version are exactly the same.  
1. The multiplexed version of the server, client can open the compressed transmission, the parameter is --c.  
1. server, client either open compression, or not open, can not open only one.  

The following tutorial uses the "multiplexed version" as an example to illustrate how to use it.  
The intranet penetration consists of three parts: client, server, and bridge; client and server actively connect to the bridge for bridging.  

### 4.2 TCP common usage  
Background:  
- Company Machine A provides web service port 80  
- There is a VPS, public network IP: 22.22.22.22  

Demand:  
At home, you can access the port 80 of company machine A by accessing port 28080 of the VPS.  
    
Steps:  
Execute on vps  
    `./proxy bridge -p ":33080" -C proxy.crt -K proxy.key`  
    `./proxy server -r ":28080@:80" -P "127.0.0.1:33080" -C proxy.crt -K proxy.key`  
    
1. Execute on company machine A  
    `./proxy client -P "22.22.22.22:33080" -C proxy.crt -K proxy.key`  

Complete  
    
### 4.3 WeChat interface local development  
Background:  
- Your own notebook provides nginx service port 80  
- There is a VPS, public network IP: 22.22.22.22  

Demand:  
Fill in the address in the webpage callback interface configuration of WeChat's development account: http://22.22.22.22/calback.php  
Then you can access the calback.php under the 80 port of the notebook. If you need to bind the domain name, you can use your own domain name.  
For example: wx-dev.xxx.com resolves to 22.22.22.22, and then in your own notebook nginx  
Configure the domain name wx-dev.xxx.com to the specific directory.  

    
Steps:  
1. Execute on vps to ensure that port 80 of vps is not occupied by other programs.  
    `./proxy bridge -p ":33080" -C proxy.crt -K proxy.key`  
    `./proxy server -r ":80@:80" -P "22.22.22.22:33080" -C proxy.crt -K proxy.key`  

1. Execute on your laptop  
    `./proxy client -P "22.22.22.22:33080" -C proxy.crt -K proxy.key`  

Complete  
    
### 4.4 UDP common usage  
Background:  
- Company Machine A provides DNS resolution service, UDP: port 53  
- There is a VPS, public network IP: 22.22.22.22  
    
Demand:  
At home, you can use the company machine A to perform domain name resolution services by setting the local dns to 22.22.22.22.  
    
Steps:  
Execute on vps  
    `./proxy bridge -p ":33080" -C proxy.crt -K proxy.key`  
    `./proxy server --udp -r ":53@:53" -P "127.0.0.1:33080" -C proxy.crt -K proxy.key`  

1. Execute on company machine A  
    `./proxy client -P "22.22.22.22:33080" -C proxy.crt -K proxy.key`  

Complete  
    
### 4.5 advanced usage one  
Background:  
- Company Machine A provides web service port 80  
- There is a VPS, public network IP: 22.22.22.22  
    
Demand:  
In order to be safe, I don't want to have access to the company machine A on the VPS, and I can access the port 28080 of the machine at home.  
Access to port 80 of company machine A via an encrypted tunnel.  
    
Steps:  
Execute on vps  
    `./proxy bridge -p ":33080" -C proxy.crt -K proxy.key`  
    
1. Execute on company machine A  
    `./proxy client -P "22.22.22.22:33080" -C proxy.crt -K proxy.key`  
    
1. Execute on your home computer  
    `./proxy server -r ":28080@:80" -P "22.22.22.22:33080" -C proxy.crt -K proxy.key`  
    
Complete  
    
### 4.6 Advanced Usage II  
Tip:  
If multiple clients are connected to the same bridge at the same time, you need to specify a different key, which can be set by the --k parameter, and --k can be any unique string.  
Just be the only one on the same bridge.  
When the server is connected to the bridge, if there are multiple clients connecting to the same bridge at the same time, you need to use the --k parameter to select the client.  
Expose multiple ports by repeating the -r parameter. The format of -r is: "local IP: local port @clientHOST:client port".  
    
Background:  
- Company Machine A provides web service port 80, ftp service port 21  
- There is a VPS, public network IP: 22.22.22.22  
    
Demand:  
At home, you can access the port 80 of company machine A by accessing port 28080 of the VPS.  
At home, I can access the 21 port of company machine A by accessing port 29090 of the VPS.  
    
Steps:  
Execute on vps  
    `./proxy bridge -p ":33080" -C proxy.crt -K proxy.key`  
    `./proxy server -r ":28080@:80" -r ":29090@:21" --k test -P "127.0.0.1:33080" -C proxy.crt -K proxy.key`  

1. Execute on company machine A  
    `./proxy client --k test -P "22.22.22.22:33080" -C proxy.crt -K proxy.key`  

Complete  
    
### 4.7.server -r parameter  
  The full format of -r is: `PROTOCOL://LOCAL_IP:LOCAL_PORT@[CLIENT_KEY]CLIENT_LOCAL_HOST:CLIENT_LOCAL_PORT`  
    
  4.7.1. Protocol PROTOCOL: tcp or udp.  
  For example: `-r "udp://:10053@:53" -r "tcp://:10800@:1080" -r ":8080@:80"`  
  If the --udp parameter is specified, PROTOCOL defaults to udp, then:`-r ":8080@:80"` defaults to udp;  
  If the --udp parameter is not specified, PROTOCOL defaults to tcp, then: `-r ":8080@:80"` defaults to tcp;  
    
  4.7.2. CLIENT_KEY: The default is default.  
  For example: -r "udp://:10053@[test1]:53" -r "tcp://:10800@[test2]:1080" -r ":8080@:80"  
  If the --k parameter is specified, such as --k test, then: `-r ":8080@:80"`CLIENT_KEY defaults to test;  
  If the --k parameter is not specified, then: `-r ":8080@:80"`CLIENT_KEY defaults to default;  
    
  4.7.3. LOCAL_IP is empty. The default is: `0.0.0.0`, CLIENT_LOCAL_HOST is empty. The default is: `127.0.0.1`;  

### 4.8.server and client connect bridge through proxy  
Sometimes the network where the server or client is located cannot directly access the external network. You need to use an https or socks5 proxy to access the Internet. Then this time  
The -J parameter can help you to connect the server or client to the bridge via https or socks5.  
The -J parameter format is as follows:  

Https proxy writing:  
The proxy needs authentication, username: username password: password  
Https://username:password@host:port  
Agent does not require authentication  
Https://host:port  

Socks5 proxy writing:  
The proxy needs authentication, username: username password: password  
Socks5://username:password@host:port  
Agent does not require authentication  
Socks5://host:port  

Host: the IP or domain name of the proxy  
Port: the port of the proxy  

### 4.9. Intranet penetration HTTP service  

Usually the HTTP request client will use the server's ip and port to set the HOST field, but it is not the same as the expected backend actual HOST, which causes tcp to be passed.  
However, the backend relies on the HOST field to locate the virtual host and it will not work. Now use the --http-host parameter to force the HOST field value of the http header to be the actual value of the backend.  
Domain names and ports can be easily solved.  

The format of the `server`-http-host parameter is as follows:  

`--http-host www.test.com:80@2200`, if the server listens to multiple ports, just repeat the `--http-host` parameter to set the HOST for each port.  

Example:  

For example, the client local nginx, 127.0.0.1:80 provides a web service, which is bound to a domain name `local.com`.  

Then the server startup parameters can be as follows:  

`proxy server -P :30000 -r :2500@127.0.0.1:80 --http-host local.com@2500`  

Explanation:  

`-r :2500@127.0.0.1:80` and `--http-host local.com:80@2500` The 2500 port is the port that the server listens locally.  

When the http protocol is used to request the ip:2500 port of the server, the header HOST field of http will be set to `local.com`.  

### 4.10 About traffic statistics  
If you start a server docking peer separately, it is the proxy-admin control panel. You need to create a new mapping in the upper-level control panel to obtain the ID of the mapping rule.  

Then start the server and add the parameter --server-id=the ID of the mapping rule to count the traffic.  

### 4.11 About p2p  
Intranet penetration support When the server and client network conditions are met, the server and client are directly connected through p2p. The opening method is:  

When starting the bridge, server, client, add the `--p2p` parameter. The server's -r parameter can be used to enable p2p (ptcp and pudp) for the port.  

If the p2p hole fails between the server and the client, the bridge transfer data is automatically switched.  

### 4.12 Client key whitelist  
The intranet penetrating bridge can set the client key whitelist. The parameter is --client-keys. The format can be:  

a. File name, file content One client key can only contain the alphanumeric underscore, which is the value of the client startup parameter --k. Only the client key can connect to the whitelist client. The line starting with # is a comment.  

b. The base64 encoding at the beginning of "base64://" is the content of the file described in a above, for example: base64://ajfpoajsdfa=  

c. "str://" multiple keywords separated by a comma at the beginning, such as: str://default,company,school  

The default is empty, allowing all keys.  

### 4.13 Network NAT Type Judgment  

Senat type judgment, easy to check whether the network supports p2p, you can execute: `proxy tools -a nattype`  

### 4.14 View help  
`./proxy help bridge`  
`./proxy help server`  
`./proxy help client`  

## 5.SOCKS5 Proxies  
prompt:  

SOCKS5 proxy, support CONNECT, UDP protocol, does not support BIND, supports username and password authentication.  

*** If your VPS is Alibaba Cloud, Tencent Cloud is a VPS, if ifconfig can't see your public IP, you can only see the intranet IP, ***  

*** Then you need to add the `-g VPS public network IP` parameter, the UDP function of the SOCKS5 proxy can work normally. ***  

### 5.1. Ordinary SOCKS5 Agent  
`./proxy socks -t tcp -p "0.0.0.0:38080"`  
    
Listen port argument `-p` can be:

```text
  -p ":8081"  listen on 8081
  -p ":8081,:8082"  listen on 8081 and 8082
  -p ":8081,:8082,:9000-9999" listen on 8081 and 8082 and 9000 and 9001 to 9999, 1002 total ports  
```

### 5.2. Ordinary secondary SOCKS5 agent  
![5.2](https://raw.githubusercontent.com/snail007/goproxy/master/doc/images/socks-2.png)  
Use local port 8090, assuming the upstream SOCKS5 proxy is `22.22.22.22:8080`  
`./proxy socks -t tcp -p "0.0.0.0:8090" -T tcp -P "22.22.22.22:8080" `  
We can also specify the black and white list file of the website domain name, one domain name and one domain name, the matching rule is the rightmost match, for example: baidu.com, the match is *.*.baidu.com, the blacklist domain name domain name goes directly to the upstream agent, white The domain name of the list does not go to the upstream agent; if the domain name is in the blacklist and in the whitelist, the blacklist works.  
`./proxy socks -p "0.0.0.0:8090" -T tcp -P "22.22.22.22:8080" -b blocked.txt -d direct.txt`  
    
### 5.3. SOCKS Level 2 Agent (Encryption)  
![5.3](https://raw.githubusercontent.com/snail007/goproxy/master/doc/images/socks-tls-2.png)  
Level 1 SOCKS proxy (VPS, IP: 22.22.22.22)  
`./proxy socks -t tls -p ":38080" -C proxy.crt -K proxy.key`  
    
Secondary SOCKS proxy (local Linux)  
`./proxy socks -t tcp -p ":8080" -T tls -P "22.22.22.22:38080" -C proxy.crt -K proxy.key`  
Then access the local port 8080 is to access the proxy port 38080 on the VPS.  
    
Secondary SOCKS proxy (local windows)  
`./proxy.exe socks -t tcp -p ":8080" -T tls -P "22.22.22.22:38080" -C proxy.crt -K proxy.key`  
Then set your windos system, the proxy that needs to go through the proxy Internet program is the socks5 mode, the address is: 127.0.0.1, the port is: 8080, the program can access the Internet through vps through the encrypted channel.  
    
### 5.4. SOCKS Level 3 Agent (Encryption)  
![5.4](https://raw.githubusercontent.com/snail007/goproxy/master/doc/images/socks-tls-3.png)  
Level 1 SOCKS proxy VPS_01, IP: 22.22.22.22  
`./proxy socks -t tls -p ":38080" -C proxy.crt -K proxy.key`  
Secondary SOCKS proxy VPS_02, IP: 33.33.33.33  
`./proxy socks -t tls -p ":28080" -T tls -P "22.22.22.22:38080" -C proxy.crt -K proxy.key`  
Level 3 SOCKS proxy (local)  
`./proxy socks -t tcp -p ":8080" -T tls -P "33.33.33.33:28080" -C proxy.crt -K proxy.key`  
Then accessing the local port 8080 is to access the proxy port 38080 on the first-level SOCKS proxy.  
    
### 5.5. SOCKS proxy traffic is forced to go to the upper level SOCKS proxy  
By default, the proxy will intelligently determine whether a website domain name is inaccessible. If it is not accessible, it will go to the upstream SOCKS proxy. With --always, all SOCKS proxy traffic can be forced to go to the upper SOCKS proxy.  
`./proxy socks --always -t tls -p ":28080" -T tls -P "22.22.22.22:38080" -C proxy.crt -K proxy.key`  
    
### 5.6. SOCKS via SSH relay  
![5.6](https://raw.githubusercontent.com/snail007/goproxy/master/doc/images/socks-ssh.png)  
Description: The principle of ssh transfer is to use the forwarding function of ssh, that is, after you connect to ssh, you can access the target address through ssh proxy.  
Suppose there is: vps  
- IP is 2.2.2.2, ssh port is 22, ssh username is: user, ssh user password is: demo  
- The user's ssh private key name is user.key  

#### *5.6.1 How to ssh username and password*  
Local SOCKS5 proxy port 28080, execute:  
`./proxy socks -T ssh -P "2.2.2.2:22" -u user -D demo -t tcp -p ":28080"`  
#### *5.6.2 How to ssh username and key*  
Local SOCKS5 proxy port 28080, execute:  
`./proxy socks -T ssh -P "2.2.2.2:22" -u user -S user.key -t tcp -p ":28080"`  

Then access the local port 28080 is to access the target address through the VPS.  

### 5.7. Certification  
For the socks5 proxy protocol, we can perform username and password authentication. The authenticated username and password can be specified on the command line.  
`./proxy socks -t tcp -p ":33080" -a "user1:pass1" -a "user2:pass2"`  
For multiple users, repeat the -a parameter.  
It can also be placed in a file in the format of a "username:password" and then specified with -F.  
`./proxy socks -t tcp -p ":33080" -F auth-file.txt`  

In addition, the socks5 agent also integrates external HTTP API authentication. We can specify an http url interface address with the --auth-url parameter.  
Then when there is a user connection, the proxy will request the url in GET mode, with the following three parameters. If the HTTP status code 204 is returned, the authentication is successful.  
In other cases, the authentication failed.  
For example:  
`./proxy socks -t tcp -p ":33080" --auth-url "http://test.com/auth.php"`  
When the user connects, the proxy will request the url ("http://test.com/auth.php") in GET mode.  
Bring four parameters: user, pass, ip, local_ip:  
Http://test.com/auth.php?user={USER}&pass={PASS}&ip={IP}&local_ip={LOCAL_IP}  
User: username  
Pass: password  
Ip: User's IP, for example: 192.168.1.200  
Local_ip: IP of the server accessed by the user, for example: 3.3.3.3  

If there is no -a or -F or --auth-url parameter, the authentication is turned off.  

### 5.8.KCP protocol transmission  
The KCP protocol requires the --kcp-key parameter to set a password for encrypting and decrypting data.  

Level 1 HTTP proxy (VPS, IP: 22.22.22.22)  
`./proxy socks -t kcp -p ":38080" --kcp-key mypassword`  
    
Secondary HTTP proxy (local Linux)  
`./proxy socks -t tcp -p ":8080" -T kcp -P "22.22.22.22:38080" --kcp-key mypassword`  
Then access the local port 8080 is to access the proxy port 38080 on the VPS, the data is transmitted through the kcp protocol.  

### 5.9. Custom DNS  
--dns-address and --dns-ttl parameters, used to specify the dns (--dns-address) used by the proxy to access the domain name.  
And the analysis result cache time (--dns-ttl) seconds, to avoid system dns interference to the proxy, in addition to the cache function can also reduce the dns resolution time to improve access speed.  
For example:  
`./proxy socks -p ":33080" --dns-address "8.8.8.8:53" --dns-ttl 300`  

### 5.10 Custom Encryption  
The proxy's socks proxy can encrypt tcp data through tls standard encryption and kcp protocol on top of tcp. In addition, it supports custom encryption after tls and kcp, which means that custom encryption and tls|kcp can be used together. The internal use of AES256 encryption, you only need to define a password when you use it.  
Encryption is divided into two parts, one is whether the local (-z) encryption and decryption, and the other is whether the transmission with the upstream (-Z) is encrypted or decrypted.  

Custom encryption requires both sides to be proxy.  

The following two levels, three levels for example:  

Secondary instance  
Execute on level 1 vps (ip: 2.2.2.2):  
`proxy socks -t tcp -z demo_password -p :7777`  
Local secondary execution:  
`proxy socks -T tcp -P 2.2.2.2:777 -Z demo_password -t tcp -p :8080`  
In this way, when the website is accessed through the local agent 8080, the target website is accessed through encrypted transmission with the upstream.  


Three-level instance  
Execute on level 1 vps (ip: 2.2.2.2):  
`proxy socks -t tcp -z demo_password -p :7777`  
Execute on the secondary vps (ip: 3.3.3.3):  
`proxy socks -T tcp -P 2.2.2.2:7777 -Z demo_password -t tcp -z other_password -p :8888`  
Local three-level execution:  
`proxy socks -T tcp -P 3.3.3.3:8888 -Z other_password -t tcp -p :8080`  
In this way, when the website is accessed through the local agent 8080, the target website is accessed through encrypted transmission with the upstream.  

### 5.11 Compressed transmission  
The proxy's socks proxy can encrypt tcp data through custom encryption and tls standard encryption and kcp protocol on top of tcp. It can also be used before custom encryption.  
Compress the data, that is, the compression function and the custom encryption and tls|kcp can be used in combination, and the compression is divided into two parts.  
Part of it is local (-m) compression transmission, and part is whether the transmission with the upstream (-M) is compressed.  

Compression requires both sides to be proxy, and compression also protects (encrypts) data to some extent.  

The following two levels, three levels for example:  

Secondary instance  

Execute on level 1 vps (ip: 2.2.2.2):  
`proxy socks -t tcp -m -p :7777`  
Local secondary execution:  
`proxy socks -T tcp -P 2.2.2.2:777 -M -t tcp -p :8080`  
In this way, when the website is accessed through the local agent 8080, the target website is accessed through compression with the upstream.  


Three-level instance  

Execute on level 1 vps (ip: 2.2.2.2):  
`proxy socks -t tcp -m -p :7777`  
Execute on the secondary vps (ip: 3.3.3.3):  
`proxy socks -T tcp -P 2.2.2.2:7777 -M -t tcp -m -p :8888`  
Local three-level execution:  
`proxy socks -T tcp -P 3.3.3.3:8888 -M -t tcp -p :8080`  
In this way, when the website is accessed through the local agent 8080, the target website is accessed through compression with the upstream.  


### 5.12 Load Balancing  

The SOCKS proxy supports the upper-level load balancing, and multiple upstream repeat-P parameters can be used.  

`proxy socks --lb-method=hash -T tcp -P 1.1.1.1:33080 -P 2.1.1.1:33080 -P 3.1.1.1:33080 -p :33080 -t tcp`  

### 5.12.1 Setting the retry interval and timeout time  

`proxy socks --lb-method=leastconn --lb-retrytime 300 --lb-timeout 300 -T tcp -P 1.1.1.1:33080 -P 2.1.1.1:33080 -P 3.1.1.1:33080 -p :33080 -t tcp`  

### 5.12.2 Setting weights  

`proxy socks --lb-method=weight -T tcp -P 1.1.1.1:33080?w=1 -P 2.1.1.1:33080?w=2 -P 3.1.1.1:33080?w=1 -p :33080 -t tcp`  

### 5.12.3 Use the target address to select the upstream  

`proxy socks --lb-hashtarget --lb-method=hash -T tcp -P 1.1.1.1:33080 -P 2.1.1.1:33080 -P 3.1.1.1:33080 -p :33080 -t tcp`  

### 5.13 Speed ​​limit  

The speed limit is 100K, which can be specified by the `-l` parameter, for example: 100K 2000K 1M . 0 means no limit.  

`proxy socks -t tcp -p 2.2.2.2:33080 -l 100K`  

### 5.14 Specifying Outgoing IP  

The `--bind-listen` parameter can be used to open the client connection with the portal IP, and use the portal IP as the outgoing IP to access the target website. If the ingress IP is an intranet IP, the egress IP does not use the ingress IP.  

`proxy socks -t tcp -p 2.2.2.2:33080 --bind-listen`  

### 5.15 Cascade Certification  

SOCKS5 supports cascading authentication, and -A can set upstream authentication information.  

upstream:  

`proxy socks -t tcp -p 2.2.2.2:33080 -a user:pass`  

local:  

`proxy socks -T tcp -P 2.2.2.2:33080 -A user:pass -t tcp -p :33080`  

### 5.16 Certificate parameters use base64 data  

By default, the -C, -K parameter is the path to the crt certificate and the key file.  

If it is the beginning of base64://, then the latter data is considered to be base64 encoded and will be used after decoding.  


### 5.17 Intelligent mode  
Intelligent mode setting, can be one of intelligent|direct|parent.  
The default is: intelligent.  
The meaning of each value is as follows:  
`--intelligent=direct`, the targets in the blocked are not directly connected.  
`--intelligent=parent`, the target that is not in the direct is going to the higher level.  
`--intelligent=intelligent`, blocked and direct have no targets, intelligently determine whether to use the upstream access target.  

### 5.18. View help  
`./proxy help socks`  

## 6.SPS Protocol Convert  

### 6.1 Function introduction  
The proxy protocol conversion uses the sps subcommand. The sps itself does not provide the proxy function. It only accepts the proxy request to "convert and forward" to the existing http(s) proxy or the socks5 proxy or ss proxy; the sps can put the existing http (s s) The proxy or socks5 proxy or ss proxy is converted to a port that supports both http(s) and socks5 and ss proxies, and the http(s) proxy supports forward proxy and reverse proxy (SNI), converted SOCKS5 proxy, UDP function is still supported when the upper level is SOCKS5 or SS; in addition, for the existing http(s) proxy or socks5 proxy, three modes of tls, tcp, and kcp are supported, and chain connection is supported, that is, multiple sps node levels can be supported. The connection builds an encrypted channel.  

The encryption methods supported by the `ss` function are: aes-128-cfb, aes-128-ctr, aes-128-gcm, aes-192-cfb, aes-192-ctr, aes-192-gcm, aes-256- Cfb , aes-256-ctr , aes-256-gcm , bf-cfb , cast5-cfb , chacha20 , chacha20-ietf , chacha20-ietf-poly1305 , des-cfb , rc4-md5 , rc4-md5-6 , salsa20 , Xchacha20  

Listen port argument `-p` can be:

```text
  -p ":8081"  listen on 8081
  -p ":8081,:8082"  listen on 8081 and 8082
  -p ":8081,:8082,:9000-9999" listen on 8081 and 8082 and 9000 and 9001 to 9999, 1002 total ports  
```

### 6.2 HTTP(S) to HTTP(S)+SOCKS5+SS  
Suppose there is already a normal http(s) proxy: 127.0.0.1:8080. Now we turn it into a common proxy that supports both http(s) and socks5 and ss. The converted local port is 18080, ss encryption: Aes-192-cfb, ss password: pass.  
The command is as follows:  
`./proxy sps -S http -T tcp -P 127.0.0.1:8080 -t tcp -p :18080 -h aes-192-cfb -j pass`  

Suppose there is already a tls http(s) proxy: 127.0.0.1:8080. Now we turn it into a normal proxy that supports both http(s) and socks5 and ss. The converted local port is 18080, and tls requires a certificate file. , ss encryption: aes-192-cfb, ss password: pass.  
The command is as follows:  
`./proxy sps -S http -T tls -P 127.0.0.1:8080 -t tcp -p :18080 -C proxy.crt -K proxy.key -h aes-192-cfb -j pass`  

Suppose there is already a kcp http(s) proxy (password is: demo123): 127.0.0.1:8080, now we turn it into a normal proxy that supports both http(s) and socks5 and ss. The converted local port is 18080, ss encryption: aes-192-cfb, ss password: pass.  
The command is as follows:  
`./proxy sps -S http -T kcp -P 127.0.0.1:8080 -t tcp -p :18080 --kcp-key demo123 -h aes-192-cfb -j pass`  

### 6.3 SOCKS5 to HTTP(S)+SOCKS5+SS  
Suppose there is already a normal socks5 proxy: 127.0.0.1:8080, now we turn it into a common proxy that supports both http(s) and socks5 and ss. The converted local port is 18080, ss encryption: aes-192 -cfb, ss password: pass.  
The command is as follows:  
`./proxy sps -S socks -T tcp -P 127.0.0.1:8080 -t tcp -p :18080 -h aes-192-cfb -j pass`  

Suppose there is already a tls socks5 proxy: 127.0.0.1:8080, now we turn it into a common proxy that supports both http(s) and socks5 and ss. The converted local port is 18080, tls requires certificate file, ss encryption Mode: aes-192-cfb, ss password: pass.  
The command is as follows:  
`./proxy sps -S socks -T tls -P 127.0.0.1:8080 -t tcp -p :18080 -C proxy.crt -K proxy.key -h aes-192-cfb -j pass`  

Suppose there is already a kcp socks5 proxy (password: demo123): 127.0.0.1:8080, now we turn it into a common proxy that supports both http(s) and socks5 and ss. The converted local port is 18080, ss Encryption method: aes-192-cfb, ss password: pass.  
The command is as follows:  
`./proxy sps -S socks -T kcp -P 127.0.0.1:8080 -t tcp -p :18080 --kcp-key demo123 -h aes-192-cfb -j pass`  

### 6.4 SS to HTTP(S)+SOCKS5+SS  
SPS upstream and local support ss protocol, the upstream can be SPS or standard ss service.  
SPS locally provides HTTP(S)\SOCKS5\SPS three defaults. When the upstream is SOCKS5, the converted SOCKS5 and SS support UDP.  
Suppose there is already a normal SS or SPS proxy (ss is enabled, encryption: aes-256-cfb, password: demo): 127.0.0.1:8080, now we turn it to support both http(s) and socks5 and The ordinary proxy of ss, the converted local port is 18080, the converted ss encryption mode: aes-192-cfb, ss password: pass.  
The command is as follows:  
`./proxy sps -S ss -H aes-256-cfb -J pass -T tcp -P 127.0.0.1:8080 -t tcp -p :18080 -h aes-192-cfb -j pass`.  

### 6.5 Chained connection  
![6.4](https://raw.githubusercontent.com/snail007/goproxy/master/doc/images/sps-tls.png)  
The above mentioned multiple sps nodes can be connected to build encrypted channels in a hierarchical connection, assuming the following vps and the home PC.  
Vps01:2.2.2.2  
Vps02:3.3.3.3  
Now we want to use pc and vps01 and vps02 to build an encrypted channel. This example uses tls encryption or kcp. Accessing local 18080 port on the PC is to access the local 8080 port of vps01.  
First on vps01 (2.2.2.2) we run a locally accessible http(s) proxy and execute:  
`./proxy http -t tcp -p 127.0.0.1:8080`  

Then run a sps node on vps01 (2.2.2.2) and execute:  
`./proxy sps -S http -T tcp -P 127.0.0.1:8080 -t tls -p :8081 -C proxy.crt -K proxy.key`  

Then run a sps node on vps02 (3.3.3.3) and execute:  
`./proxy sps -S http -T tls -P 2.2.2.2:8081 -t tls -p :8082 -C proxy.crt -K proxy.key`  

Then run a sps node on the pc and execute:  
`./proxy sps -S http -T tls -P 3.3.3.3:8082 -t tcp -p :18080 -C proxy.crt -K proxy.key`  

carry out.  

### 6.6 Listening to multiple ports  
In general, listening to a port is fine, but if you need to listen to both ports 80 and 443 as a reverse proxy, the -p parameter is supported.  
The format is: `-p 0.0.0.0:80, 0.0.0.0:443`, multiple bindings can be separated by commas.  

### 6.7 Authentication  
Sps supports http(s)\socks5 proxy authentication, which can be cascaded and has four important pieces of information:  
1: The user sends the authentication information `user-auth`.  
2: Set the local authentication information `local-auth`.  
3: Set the connection authentication information 'parent-auth` used by the upstream.  
4: The authentication information `auth-info-to-parent` that is finally sent to the upstream.  
Their situation is as follows:  

User-auth | local-auth | parent-auth | auth-info-to-paren  
| ------ | ------ | ------ | ------  
| Yes / No | Yes | Yes | From parent-auth  
| Yes / No | No | Yes | From parent-auth  
| Yes / No | Yes | No | No  
| No | No | No | No  
| Yes | No | No | From user-auth  

For the sps proxy we can perform username and password authentication. The authenticated username and password can be specified on the command line.  
`./proxy sps -S http -T tcp -P 127.0.0.1:8080 -t tcp -p ":33080" -a "user1:pass1:0:0:" -a "user2:pass2:0:0: "`  
For multiple users, repeat the -a parameter.  
Can also be placed in a file, the format is one line a `username: password: number of connections: rate: upstream`, and then specified with -F.  
`./proxy sps -S http -T tcp -P 127.0.0.1:8080 -t tcp -p ":33080" -F auth-file.txt`  

If the upstream has authentication, the lower level can set the authentication information with the -A parameter, for example:  
upstream: `./proxy sps -S http -T tcp -P 127.0.0.1:8080 -t tcp -p ":33080" -a "user1:pass1:0:0:" -a "user2:pass2:0: 0:"`  
Subordinate: `./proxy sps -S http -T tcp -P 127.0.0.1:8080 -A "user1:pass1" -t tcp -p ":33080" `  

For more details on certification, please refer to `9.API Certification` and `10.Local Certification`  

### 6.8 Multiple Upstream  

If there are multiple upstreams, they can be specified by multiple -Ps.  

such as:  

`proxy sps -P http://127.0.0.1:3100 -P socks5://127.0.0.1:3200`  

The complete format of `-P` is as follows:  

 `protocol://a:b@2.2.2.2:33080#1`  

Each section is explained below:  

`protocol://` is the protocol type, possible types and contains the following:  

```text  
Http is equivalent to -S http -T tcp  
Https is equivalent to -S http -T tls --parent-tls-single , which is http(s) proxy over TLS  
Https2 is equivalent to -S http -T tls  
Socks5 is equivalent to -S socks -T tcp  
Socks5s is equivalent to -S socks -T tls --parent-tls-single , which is socks over TLS  
Socks5s2 is equivalent to -S socks -T tls  
Ss is equivalent to -S ss -T tcp  
Httpws is equivalent to -S http -T ws  
Httpwss is equivalent to -S http -T wss  
Socks5ws is equivalent to -S socks -T ws  
Socks5wss is equivalent to -S socks -T wss  
```  

`a:b` is the username and password of the proxy authentication. If it is ss, `a` is the encryption method, `b` is the password, and no username password can be left blank, for example: `http://2.2.2.2:33080` If the username and password are protected, special symbols can be encoded using urlencode.  

`2.2.2.2:33080` is the upstream address, the format is: `IP (or domain name): port `, if the underlying is ws/wss protocol can also bring the path, such as: `2.2.2.2: 33080/ws`;  
You can also set the `encryption method` and `password` of `ws\wss` by appending the query parameters `m` and `k`, for example: `2.2.2.2:33080/ws?m=aes-192-cfb&k=password`  

`#1` When multiple upper-level load balancing is a weighting strategy, the weights are rarely used.  

### 6.9 Custom Encryption  
The proxy sps proxy can encrypt tcp data through tls standard encryption and kcp protocol on top of tcp, in addition to support after tls and kcp  
Custom encryption, that is, custom encryption and tls|kcp can be used in combination, internally using AES256 encryption, only need to define it when using  
A password can be used, the encryption is divided into two parts, one part is whether the local (-z) encryption and decryption, and the part is the encryption and decryption with the upstream (-Z) transmission.  

Custom encryption requires both sides to be proxy.  

The following two levels, three levels for example:  

Suppose there is already an http(s) proxy: `6.6.6.6:6666`  

Secondary instance  

Execute on level 1 vps (ip: 2.2.2.2):  
`proxy sps -S http -T tcp -P 6.6.6.6:6666 -t tcp -z demo_password -p :7777`  
Local secondary execution:  
`proxy sps -T tcp -P 2.2.2.2:777 -Z demo_password -t tcp -p :8080`  
In this way, when the website is accessed through the local agent 8080, the target website is accessed through encrypted transmission with the upstream.  


Three-level instance  

Execute on level 1 vps (ip: 2.2.2.2):  
`proxy sps -S http -T tcp -P 6.6.6.6:6666 -t tcp -z demo_password -p :7777`  
Execute on the secondary vps (ip: 3.3.3.3):  
`proxy sps -T tcp -P 2.2.2.2:7777 -Z demo_password -t tcp -z other_password -p :8888`  
Local three-level execution:  
`proxy sps -T tcp -P 3.3.3.3:8888 -Z other_password -t tcp -p :8080`  
In this way, when the website is accessed through the local agent 8080, the target website is accessed through encrypted transmission with the upstream.  

### 6.10 Compressed transmission  
The proxy sps proxy can encrypt tcp data through custom encryption and tls standard encryption and kcp protocol on top of tcp. It can also be used before custom encryption.  
Compress the data, that is, the compression function and the custom encryption and tls|kcp can be used in combination, and the compression is divided into two parts.  
Part of it is local (-m) compression transmission, and part is whether the transmission with the upstream (-M) is compressed.  

Compression requires both sides to be proxy, and compression also protects (encrypts) data to some extent.  

The following two levels, three levels for example:  

Secondary instance  

Execute on level 1 vps (ip: 2.2.2.2):  
`proxy sps -t tcp -m -p :7777`  
Local secondary execution:  
`proxy sps -T tcp -P 2.2.2.2:777 -M -t tcp -p :8080`  
In this way, when the website is accessed through the local agent 8080, the target website is accessed through compression with the upstream.  


Three-level instance  

Execute on level 1 vps (ip: 2.2.2.2):  
`proxy sps -t tcp -m -p :7777`  
Execute on the secondary vps (ip: 3.3.3.3):  
`proxy sps -T tcp -P 2.2.2.2:7777 -M -t tcp -m -p :8888`  
Local three-level execution:  
`proxy sps -T tcp -P 3.3.3.3:8888 -M -t tcp -p :8080`  
In this way, when the website is accessed through the local agent 8080, the target website is accessed through compression with the upstream.  

### 6.11 Disabling the protocol  
By default, SPS supports http(s) and socks5 two proxy protocols. We can disable a protocol by parameter.  
For example:  
1. Disable the HTTP(S) proxy function to retain only the SOCKS5 proxy function, parameter: `--disable-http`.  
`proxy sps -T tcp -P 3.3.3.3:8888 -M -t tcp -p :8080 --disable-http`  

1. Disable the SOCKS5 proxy function to retain only the HTTP(S) proxy function, parameter: `--disable-socks`.  
`proxy sps -T tcp -P 3.3.3.3:8888 -M -t tcp -p :8080 --disable-http`  

### 6.12 Speed ​​limit  

Suppose there is a SOCKS5 upstream:  

`proxy socks -p 2.2.2.2:33080 -z password -t tcp`  

SPS lower level, speed limit 100K  

`proxy sps -S socks -P 2.2.2.2:33080 -T tcp -Z password -l 100K -t tcp -p :33080`  

It can be specified by the `-l` parameter, for example: 100K 2000K 1M . 0 means no limit.  

### 6.13 Specifying Outgoing IP  

The `--bind-listen` parameter can be used to open the client connection with the portal IP, and use the portal IP as the outgoing IP to access the target website. If the ingress IP is an intranet IP, the egress IP does not use the ingress IP.  

`proxy sps -S socks -P 2.2.2.2:33080 -T tcp -Z password -l 100K -t tcp --bind-listen -p :33080`  

### 6.14 Certificate parameters use base64 data  

By default, the -C, -K parameter is the path to the crt certificate and the key file.  

If it is the beginning of base64://, then the latter data is considered to be base64 encoded and will be used after decoding.  

### 6.15 Independent Service  
The sps function does not force a upstream to be specified. When the upstream is empty, the sps itself can complete the full proxy function. If the upstream is specified, the upstream connection target is used as before.  
The following command is to open the http(s)\ss\socks service with one click.  
`./proxy sps -p :33080`  

### 6.16 Target Redirection  
The https(s)\socks5\ss proxy function provided by the sps function, the client connects to the specified "target" through the sps proxy. This "target" is generally a website or an arbitrary tcp address.  
The website "target" is generally foo.com: 80, foo.com: 443, sps supports the use of the --rewrite parameter to specify a "target" redirection rule file, redirect the target, the client is non-perceived,  
For example, if you redirect to "target": demo.com:80 to 192.168.0.12:80, then the client visits the website demo.com, in fact, the website service provided by 192.168.0.12.  
Example of a "target" redirection rule file:  

```text  
# example  
Www.a.com:80 10.0.0.2:8080  
**.b.com:80 10.0.0.2:80  
192.168.0.11:80 10.0.0.2:8080  
```  

### 6.17 View help  

`./proxy help sps`  

## 7.KCP Configuration  

### 7.1 Configuration Introduction  
Many functions of the proxy support the kcp protocol. Any function that uses the kcp protocol supports the configuration parameters described here.  
Therefore, the KCP configuration parameters are introduced here.  

### 7.2 Detailed configuration  
There are a total of 17 KCP configuration parameters, you can not set them, they have default values, if for the best effect,  
You need to configure the parameters according to your own network conditions. Because the kcp configuration is complex, it requires a certain network basics.  
If you want to get more detailed configuration and explanation of kcp parameters, please search for yourself. The command line name for each parameter, along with the default values ​​and simple function descriptions are as follows:  
```  
--kcp-key="secrect" pre-shared secret between client and server  
--kcp-method="aes" encrypt/decrypt method, can be: aes, aes-128, aes-192, salsa20, blowfish,  
                           Twofish, cast5, 3des, tea, xtea, xor, sm4, none  
--kcp-mode="fast" profiles: fast3, fast2, fast, normal, manual  
--kcp-mtu=1350 set maximum transmission unit for UDP packets  
--kcp-sndwnd=1024 set send window size(num of packets)  
--kcp-rcvwnd=1024 set receive window size(num of packets)  
--kcp-ds=10 set reed-solomon erasure coding - datashard  
--kcp-ps=3 set reed-solomon erasure coding - parityshard  
--kcp-dscp=0 set DSCP(6bit)  
--kcp-nocomp disable compression  
--kcp-acknodelay be carefull! flush ack immediately when a packet is received  
--kcp-nodelay=0 be carefull!  
--kcp-interval=50 be carefull!  
--kcp-resend=0 be carefull!  
--kcp-nc=0 be carefull! no congestion  
--kcp-sockbuf=4194304 be carefull!  
--kcp-keepalive=10 be carefull!  
```  
Tip:  
Parameters: -- four fast3, fast2, fast, normal modes in kcp-mode,  
Equivalent to setting the following four parameters:  
Normal:`--nodelay=0 --interval=40 --resend=2 --nc=1`  
Fast :`--nodelay=0 --interval=30 --resend=2 --nc=1`  
Fast2:`--nodelay=1 --interval=20 --resend=2 --nc=1`  
Fast3:`--nodelay=1 --interval=10 --resend=2 --nc=1`  

## 8. Security DNS  

### 8.1 Introduction  
DNS is known as the service provided by UDP port 53, but with the development of the network, some well-known DNS servers also support TCP mode dns query, such as Google's 8.8.8.8, the DNS anti-pollution server principle of the proxy is to start a proxy DNS proxy locally. Server, which uses TCP to perform dns query through the upstream agent. If it communicates with the upstream agent, it can perform secure and pollution-free DNS resolution. It also supports independent services, concurrent parsing, and enhanced enhanced hosts file function to support flexible concurrent parsing and forwarding.  

Dns resolution order:  
1. Use the parameter --hosts to parse.  
2. If the domain name to be resolved is not found in 1, it is parsed using the parameter --forward rule.  
3. The domain name to be resolved is not found in 1 and 2, and the default --default parsing is used. The default default behavior parameter values ​​are three: proxy, direct, and system.  
    The three parameter values ​​are explained as follows:  
    Proxy: The domain name is resolved by the dns server specified by the -q parameter.  
    Direct: Connect to the dns server specified by the -q parameter to resolve the domain name through the local network.  
    System: resolves the domain name through the system dns.  

Tip:  
The host file format specified by the --hosts parameter is the same as the system hosts file, and the domain name supports wildcards. You can refer to the hosts file.  
The parsing forwarding rule file specified by the --forward parameter can be referenced to the resolve.rules file. The domain name supports wildcards. It supports multiple dns servers for each domain name to be parsed concurrently. Whoever resolves the fastest resolution will use the resolution result.  
The -q parameter can specify multiple remote dns servers to perform concurrent parsing. Whoever resolves the fastest parsing success, the default is: 1.1.1.1, 8.8.8.8, 9.9.9.9, multiple comma-separated,  
           For example, you can also bring ports: 1.1.1.1, 8.8.8.8#53, 9.9.9.9  

If you are a standalone service, you don't need a upstream:  
Can perform:  
`proxy dns --default system -p :5353`  
Or  
`proxy dns --default direct -p :5353`  

### 8.2 Example of use  

#### 8.2.1 Normal HTTP(S) upstream agent  

Suppose there is a upstream agent: 2.2.2.2:33080  
Local execution:  
`proxy dns -S http -T tcp -P 2.2.2.2:33080 -p :53`  
Then the local UDP port 53 provides DNS resolution.  

#### 8.2.2 Ordinary SOCKS5 upstream agent  

Suppose there is a upstream agent: 2.2.2.2:33080  
Local execution:  
`proxy dns -S socks -T tcp -P 2.2.2.2:33080 -p :53`  
Then the local UDP port 53 provides DNS resolution.  

#### 8.2.3 TLS encrypted HTTP(S) upstream agent  

Suppose there is a upstream agent: 2.2.2.2:33080  
The commands executed by the upstream agent are:  
`proxy http -t tls -C proxy.crt -K proxy.key -p :33080`  
Local execution:  
`proxy dns -S http -T tls -P 2.2.2.2:33080 -C proxy.crt -K proxy.key -p :53`  
Then the local UDP port 53 provides a secure anti-pollution DNS resolution function.  

#### 8.2.4 TLS-encrypted SOCKS5 upstream agent  

Suppose there is a upstream agent: 2.2.2.2:33080  
The commands executed by the upstream agent are:  
`proxy socks -t tls -C proxy.crt -K proxy.key -p :33080`  
Local execution:  
`proxy dns -S socks -T tls -P 2.2.2.2:33080 -C proxy.crt -K proxy.key -p :53`  
Then the local UDP port 53 provides a secure anti-pollution DNS resolution function.  

#### 8.2.5 KCP encrypted HTTP(S) upstream agent  

Suppose there is a upstream agent: 2.2.2.2:33080  
The commands executed by the upstream agent are:  
`proxy http -t kcp -p :33080`  
Local execution:  
`proxy dns -S http -T kcp -P 2.2.2.2:33080 -p :53`  
Then the local UDP port 53 provides a secure anti-pollution DNS resolution function.  

#### 8.2.6 KCP encrypted SOCKS5 upstream agent  

Suppose there is a upstream agent: 2.2.2.2:33080  
The commands executed by the upstream agent are:  
`proxy socks -t kcp -p :33080`  
Local execution:  
`proxy dns -S socks -T kcp -P 2.2.2.2:33080 -p :53`  
Then the local UDP port 53 provides a secure anti-pollution DNS resolution function.  

#### 8.2.7 Custom encrypted HTTP(S) upstream agent  

Suppose there is a upstream agent: 2.2.2.2:33080  
The commands executed by the upstream agent are:  
`proxy http -t tcp -p :33080 -z password`  
Local execution:  
`proxy dns -S http -T tcp -Z password -P 2.2.2.2:33080 -p :53`  
Then the local UDP port 53 provides a secure anti-pollution DNS resolution function.  

#### 8.2.8 Custom encrypted SOCKS5 upstream agent  
   
Suppose there is a upstream agent: 2.2.2.2:33080  
The commands executed by the upstream agent are:  
`proxy socks -t kcp -p :33080 -z password`  
Local execution:  
`proxy dns -S socks -T tcp -Z password -P 2.2.2.2:33080 -p :53`  
Then the local UDP port 53 provides a secure anti-pollution DNS resolution function.  

## 9.API Authentication  

The proxy's http(s)/socks5/sps proxy function supports user-to-agent access via the API.  

### What can I do through the API?  

- User dimension, which controls the single connection rate and controls the maximum number of connections.  
- IP dimension, which controls the single connection rate and controls the maximum number of connections.  
- Dynamic upstream, can dynamically obtain its upstream from the API according to the user or client IP, and support http(s)/socks5/ss upstream.  
- Authenticate every connection, regardless of whether client authentication is required.  
- Cache authentication results, time can be set to reduce API pressure.  

#### Specific use  
The proxy's http(s)/socks5/sps proxy API function is controlled by three parameters: `--auth-url` and `--auth-nouser` and `--auth-cache`.  
The parameter `--auth-url` is the HTTP API interface address. When the client connects, the proxy will request the url in GET mode, with the following parameters. If the HTTP status code 204 is returned, the authentication is successful. In other cases, the authentication fails.  

An example of a complete request API:  
`http://test.com/auth.php?user=a&pass=b&client_addr=127.0.0.1:49892&local_addr=127.0.0.1:8100&target=http%3A%2F%2Fwww.baidu.com&service=http&sps=0`  

#### Parameter Description  
`user and pass` When the proxy turns on authentication, here is the username and password provided by the client.  
`client_addr` The address used by the client to access the proxy, format IP: port.  
`local_addr` The proxy address accessed by the client, format IP: port.  
`service` Proxy type, divided into: http, socks.  
Whether the `sps` proxy is provided by sps, 1: yes, 0: no.  
`target` The target to be accessed by the client. If it is an http(s) proxy, the target is the specific url accessed; if it is a socks5 proxy, the target is empty.  

#### Example  
Suppose --auth-url http://127.0.0.1:333/auth.php points to a php interface address.  
The contents of auth.php are as follows:  

```php  
<?php  
$proxy_ip=$_GET['local_addr'];  
$user_ip=$_GET['client_addr'];  
$service=$_GET['service'];  
$is_sps=$_GET['sps']=='1';  
$user=$_GET['user'];  
$pass=$_GET['pass'];  
$target=$_GET['target'];  
/ / Business logic judgment  
//....  

/ / Set the authentication result  
Header("userconns:1000");  
Header("ipconns:2000");  
Header("userrate:3000");  
Header("iprate:8000");  
Header("UPSTREAM:http://127.0.0.1:3500?parent-type=tcp");  
Header("HTTP/1.1 204 No Content");  
```  

#### Explanation  
Userconns: The maximum number of connections for the user, not limited to 0 or not set this header.  
Ipcons: The maximum number of connections for the user IP, not limited to 0 or not set this header.  
Userrate: User's single TCP connection rate limit, in bytes/second, is not limited to 0 or does not set this header.  
Iprate: The single TCP connection rate limit of the user IP, in bytes/second, not limited to 0 or not set this header.  
Upstream: The upstream used, not empty, or not set this header.  

#### Tips  
1. By default, `--auth-url` is required to provide the user name and password. If you do not need the client to provide the username and password, and authenticate, you can add `--auth-nouser`. The visit will still access the authentication address `--auth-url` for authentication. Only the $user authentication username and the $pass authentication password received in the php interface are empty.  
2. Connection limit priority: User authentication file rate limit - "File ip.limit rate limit -" API user rate limit - "API IP rate limit -" command line global connection limit.  
3. Rate Limit Priority: User Authentication File Rate Limit - "File ip.limit Rate Limit -" API User Rate Limit - "API IP Rate Limit - "Command Line Global Rate Limit.  
3. The upstream obtains the priority: the upstream of the user authentication file - the file ip.limit upstream-"API upstream-" command line specifies the upstream.  
4.`--auth-cache` authentication cache, cache the authentication result for a certain period of time, improve performance, reduce the pressure on the authentication interface, --auth-cache unit seconds, default 60, set 0 to close the cache.  

#### upstream detailed description  

1. When the parameter `sps` is 0.  
When the service is http, upstream only supports http(s) proxy, and does not support authentication. If authentication is required, it can be replaced by sps. Format:  
  `http://127.0.0.1:3100?argk=argv`  
When the service is a socks, the upstream only supports the socks5 proxy. The format is:  
  `socks://127.0.0.1:3100?argk=argv`  

Explanation: `http://`,`socks://` is fixed, `127.0.0.1:3100` is the address of the upstream  

2. When `sps` is 1.  
Upstream supports socks5, http(s) proxy, support authentication, format: `protocol://a:b@2.2.2.2:33080?argk=argv`, please refer to SPS chapter for details, **multiple upstreams** , the description of the `-P` parameter.  
3. Parameters, `?` followed by `argk=argv` are parameters: parameter name = parameter value, multiple parameters are connected with `&`.  
  All the supported parameters are as follows, and the meaning of the command line with the same name is the same.  
  1. parent-type : upper-level transport type, support tcp, tls, ws, wss  
  2. parent-ws-method: The encryption method of the upper-level ws transmission type, the supported value is the same as the value range supported by the command line.  
  3. parent-ws-password: The upper-level ws transmission type encryption password, the alphanumeric password  
  4. parent-tls-single : Whether the upper-level tls transport type is a one-way tls, which can be: true | false  
  5. timeout : timeout for establishing tcp connection, number, in milliseconds  
  6. ca : The base64-encoded string of the upper-level tls transport type ca certificate file.  
  7. cert : The base64 encoded string of the higher level tls transport type certificate file.  
  8. key : The base64 encoded string of the higher-level tls transport type certificate key file.  

## 10. Authentication  

The proxy http(s)/socks5/sps proxy function supports the user to access the proxy pair through the configuration file, and supports the http(s) proxy ``Proxy Basic proxy authentication` and the socks5 proxy authentication.  

### start using  
The proxy's http(s)/socks5/sps proxy function can pass  
`--auth-file`, `--max-conns`, `--ip-limit`, `--rate-limit`, `-a` These five parameters control.  

#### Detailed explanation of parameters  

##### `--auth-file`  
The authenticated user name and password file. This parameter specifies a file, one line per rule, in the format: "username: password: number of connections: rate: upstream".  
`Connection number` is the maximum number of connections for the user. The 'rate' is the maximum speed of each tcp connection of the user. The unit is: byte/second. The upper level is the upper level used by the user.  
Not only can the authenticated user be set by `--auth-file`, but also the `-a` parameter can be set directly. Multiple users can repeat multiple `-a` parameters.  
For example: `proxy http -a a:b:0:0: -a c:d:0:0:`  

Example explanation:  
For example: `user:pass:100:10240:http://192.168.1.1:3100`  
`user` is the authentication username  
`pass` is the authentication user password (cannot contain a colon:)  
`100` is the maximum number of connections for this user, not limited to write 0  
`10240` is the rate limit of this user's single tcp connection, the unit is: byte / sec, no limit write 0  
`http://192.168.1.1:3100` is the upstream used by this user, no space is left blank  

##### `--max-conns`  
Limit the maximum number of global connections for the proxy service, a number, 0 is unrestricted, default is 0.  

##### `--ip-limit`  
Controls the number of connections and connection rate of the client IP. This parameter specifies a file, one rule per line, and the beginning of # is gaze.  
The sample file ip.limit, the rule format is as follows:  
`127.0.0.1:100:10240:http://192.168.1.1:3100`  
Rule interpretation:  
`127.0.0.1` is the IP to be restricted  
`100` is the maximum number of connections for this IP, not limited to write 0  
`10240` is the rate limit of IP single tcp connection, the unit is: byte / s, no limit write 0  
`http://192.168.1.1:3100` is the upstream used by this IP, and it is not left blank.  

##### `--rate-limit`  
Limit the speed of each tcp connection of the service, for example: 100K 2000K 1M . 0 means unlimited, default 0.  

