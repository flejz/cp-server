#!/bin/bash

USR=""
HOST=""
LPORT=7947
LHOST=localhost
RPORT=7947
RHOST=localhost
OTHER_ARGUMENTS=""

# Loop through arguments and process them
for arg in "$@"
do
  case $arg in
    --lport=*)
      LPORT="${arg#*=}"
      shift
      ;;
    --lhost=*)
      LHOST="${arg#*=}"
      shift
      ;;
    --rport=*)
      RPORT="${arg#*=}"
      shift
      ;;
    --rhost=*)
      RHOST="${arg#*=}"
      shift
      ;;
    -h=*|--host=*)
      HOST="${arg#*=}"
      shift
      ;;
    -u=*|--user=*)
      USR="${arg#*=}"
      shift
      ;;
    *)
      OTHER_ARGUMENTS+=" $1"
      shift # Remove generic argument from processing
      ;;
  esac
done


echo $LHOST
echo $LPORT
echo $RHOST
echo $RPORT
echo $OTHER_ARGUMENTS


#ssh -f -L 7947:localhost:7947 flejz@kraken sleep 10 && telnet localhost 7947
