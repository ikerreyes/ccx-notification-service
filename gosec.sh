#!/bin/bash
# Copyright 2020, 2021, 2022 Red Hat, Inc
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

RED_BG=$(tput setab 1)
GREEN_BG=$(tput setab 2)
BLUE=$(tput setaf 4)
NC=$(tput sgr0) # No Color

GO_SEC_ARGS=""

if [[ $* != *verbose* ]]; then
    GO_SEC_ARGS="-quiet"
fi

cd "$(dirname "$0")" || exit

echo -e "${BLUE}Security issues detection${NC}"

if ! [ -x "$(command -v gosec)" ]
then
    echo -e "${BLUE}Installing ${NC}"
    curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s -- -b "$(go env GOPATH)/bin" v2.22.2
    # shellcheck disable=SC2181
    if [ $? -eq 0 ]
    then
        echo -e "${BLUE}Installed ${NC}"
    else
        echo -e "${RED_BG}[FAIL]${NC} Installation failure"
        exit 2
    fi
fi

if ! gosec $GO_SEC_ARGS ./...
then
    echo -e "${RED_BG}[FAIL]${NC} Potential security issues detected!"
    exit 1
else
    echo -e "${GREEN_BG}[OK]${NC} No potential security issues has been detected"
    exit 0
fi
