#!/usr/bin/env python
import requests
import configparser
import argparse
from urllib.parse import urljoin
from pathlib import Path

def parse_arguments():
    parser = argparse.ArgumentParser()
    parser.add_argument('--confirm', action="store_true", help="Write config to ~/.ssh/config")
    parser.add_argument('--list', action="store_true", help="Output config to stdout")
    return parser.parse_args()

def fetch_veritas_hosts(host):
    response = requests.get(urljoin(host, "/node"))
    nodes = response.json()
    names = map(lambda node: node["name"], nodes)
    return list(names)

def fetch_node_ip(host, node):
    request_path = "/node/" + node + "/prop/ip_address_0"
    response = requests.get(urljoin(host, request_path))
    data = response.json()
    ip_address = data["property_value"]
    return ip_address

def fetch_config(path):
    config = configparser.ConfigParser()
    config.read(path)
    return config

def generate_ssh_config_file(nodes, ips, user):
    content = ""
    for i in range(len(nodes)):
        content += "Host " + nodes[i] + "\n"
        content += "\tHostName " + ips[i] + "\n"
        content += "\tUser " + user + "\n\n"
    return content

def write_to_ssh_config(content):
    f = open(Path.home().joinpath(".ssh/config"), "w")
    f.write(content)
    f.close()

def main():
    args = parse_arguments()

    if not args.list: 
        print("Generating new ssh conf")

    config = fetch_config(Path.home().joinpath(".veritasrc"))
    host = config["server"]["host"]
    nodes = fetch_veritas_hosts(host)
    ips = list(map(lambda node: fetch_node_ip(host, node), nodes))
    content = generate_ssh_config_file(nodes, ips, "ansible")

    if not args.confirm or args.list:
        print(content)
        return
    
    print("Writing to ssh config")
    write_to_ssh_config(content)

if __name__ == "__main__":
    main()