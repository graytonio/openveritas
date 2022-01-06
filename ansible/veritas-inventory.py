#!/usr/bin/env python
# Gathers inventory from veritas along with all the properties with it

import os
import sys
import argparse

from ansible.plugins import inventory

try:
    import json
except ImportError:
    import simplejson

def read_cli_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('--list', action='store_true')
    parser.add_argument('--host', action = 'store')
    return parser.parse_args()

def fetch_inventory():
        return {
            'group': {
                'hosts': ['127.0.0.1'],
                'vars': {
                    'ansible_ssh_user': 'graytonio',
                    'ansible_ssh_private_key_file':
                        '~/.vagrant.d/insecure_private_key',
                    'example_variable': 'value'
                }
            },
            '_meta': {
                'hostvars': {
                    '127.0.0.1': {
                        'host_specific_var': 'foo'
                    }
                }
            }
        }

def empty_inventory(self):
    return {'_meta': {'hostvars': {}}}

def main():
    args = read_cli_args()
    inventory = {}
    if args.list:
        inventory = fetch_inventory()
    elif args.host:
        inventory = empty_inventory()
    else:
        inventory = empty_inventory()


    print(json.dumps(inventory))

    return 0

if __name__ == '__main__':
    main()