#!/bin/bash

docker run --rm -it -v ~/webapp:/webapp registry.cn-hangzhou.aliyuncs.com/deeplearns/gocentoscompiler:latest /bin/bash -c 'export PATH=$PATH:/usr/local/go/bin &&  cd /webapp && /bin/make bl'

