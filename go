#!/usr/bin/python
# gopy.py

import os
import json
import sys
import commands

def getRootPath():
    path_info = commands.getstatusoutput('cd && pwd')
    return path_info[1]

def call(data):
    if data['host'] == None:
        print "try to call invalid host"
        return
    command = 'ssh ' + data['host']
    print command
    if 'passwd' in data:
        command = "sshpass -p \"" + data['passwd'] + "\" " + command
    if 'rsa' in data:
        command = command + " -i " + data['rsa']
    if 'cmd' in data:
        command = command + " " + data['cmd']
    
    os.system(command)

def goTag(tag):
    for v in CONFIG:
        if v['tag'] == tag:
            call(v)
            return
    tag = int(tag) - 1
    if CONFIG[tag]:
        call(CONFIG[tag])
        return
    print "%s not found" % tag

def printInfo():
    i = 1
    for v in CONFIG:
        print "%s) %s\t%s" % (i, v['tag'], v['host'])
        i = i + 1
    print "q) exit"

# run

CONFIG_FILE = getRootPath() + '/.goconfig'
try:
    fileObject = open(CONFIG_FILE)
    all_content_text = fileObject.read()
except:
    print "%s not exists" % CONFIG_FILE
    exit()
finally:
    fileObject.close()
CONFIG = json.loads(all_content_text)

def start():
    printInfo()
    tag = raw_input("Enter to go: ")
    if tag == 'q' or tag == 'exit':
        exit()
    elif tag == '':
        start()
    goTag(tag)

if len(sys.argv) <= 1 :
    start()
else:
    goTag(sys.argv[1])
