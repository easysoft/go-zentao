#!/bin/bash
#
#  Copyright 2022, easysoft
#
#  Licensed under the Apache License, Version 2.0 (the "License");
#  you may not use this file except in compliance with the License.
#  You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
#  Unless required by applicable law or agreed to in writing, software
#  distributed under the License is distributed on an "AS IS" BASIS,
#  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#  See the License for the specific language governing permissions and
#  limitations under the License.
#


# curl -X POST -d '{"account": "admin", "password": "jaege1ugh4ooYip7"}'  http://127.0.0.1/api.php/v1/tokens

curl -H "Token: 1bd72495bd50e243b4fa7962023551f6" http://127.0.0.1/api.php/v1/config/version
