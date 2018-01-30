#!/bin/sh

e=`pidof $1`
sudo kill $e
