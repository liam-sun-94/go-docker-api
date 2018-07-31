# go-restfulldocker API v1.0

#### api: /BuildImage
type:POST
功能：基于提供的试题生成docker镜像
参数：body中json格式
参数示例：

>```
{
"uid":"user1",
"imgName":"img-test",
"imgId":"imgId-1",
"problem":{
	"8":{
        "name":"第8题",
        "type":"选择题",
        "problem":"1+1=？",
        "answer":"A.1&B.2&C.3&D.4@A",
        "createTime":"2018-7-6 16:34",
        "author":"slq"
    },
    "9":{
        "name":"第9题",
        "type":"选择题",
        "problem":"1+1=？",
        "answer":"A.1&B.2&C.3&D.4@A",
        "createTime":"2018-7-6 16:34",
        "author":"slq"
    }
}}
```
返回值：string（无错误返回：“BuildImage”）


#### api: /BII
type:GET
功能：用于获取BuildImage时的状态和结果
参数：

| 参数名 | 参数类型 | 示例 |
| - | :- | :- |
| img_id | string | "imgId-1" |

返回值：string 不同阶段返回值不同

#### api: /ContainerList
type:GET
功能：用于获取所有容器的信息
参数：无

返回值：json
返回值示例：此版返回信息较多，后期可以根据需求返回数据
>```
[
  {
    "Id": "725a6ee9a490cbdf02249ee0fbdd92bc2c3c29fa04b59c8631419d195eda97d2",
    "Names": [
      "/tt4"
    ],
    "Image": "alpine",
    "ImageID": "sha256:11cd0b38bc3ceb958ffb2f9bd70be3fb317ce7d255c8a4c3f4af30e298aa1aab",
    "Command": "/bin/sh",
    "Created": 1532941360,
    "Ports": [
      {
        "PrivatePort": 8080,
        "Type": "tcp"
      },
      {
        "IP": "0.0.0.0",
        "PrivatePort": 3306,
        "PublicPort": 3306,
        "Type": "tcp"
      }
    ],
    "Labels": {},
    "State": "running",
    "Status": "Up About an hour",
    "HostConfig": {
      "NetworkMode": "default"
    },
    "NetworkSettings": {
      "Networks": {
        "bridge": {
          "IPAMConfig": null,
          "Links": null,
          "Aliases": null,
          "NetworkID": "e0bccd963bea8cbf9f8efe7a61dbd7d23e896820ef29e704389d6065f061ad21",
          "EndpointID": "9442c0753b68040f204bd85f51bdff0aa49e32de10169eae10b2b6e7b977ade4",
          "Gateway": "172.17.0.1",
          "IPAddress": "172.17.0.3",
          "IPPrefixLen": 16,
          "IPv6Gateway": "",
          "GlobalIPv6Address": "",
          "GlobalIPv6PrefixLen": 0,
          "MacAddress": "02:42:ac:11:00:03",
          "DriverOpts": null
        }
      }
    },
    "Mounts": []
  },
  {
    "Id": "c1e48c1082d2b68f3f4a93a65d7658dc4ef40cebff93589dafabc6282ef46b64",
    "Names": [
      "/tt3"
    ],
    "Image": "alpine",
    "ImageID": "sha256:11cd0b38bc3ceb958ffb2f9bd70be3fb317ce7d255c8a4c3f4af30e298aa1aab",
    "Command": "/bin/sh",
    "Created": 1532940431,
    "Ports": [],
    "Labels": {},
    "State": "running",
    "Status": "Up About an hour",
    "HostConfig": {
      "NetworkMode": "default"
    },
    "NetworkSettings": {
      "Networks": {
        "bridge": {
          "IPAMConfig": null,
          "Links": null,
          "Aliases": null,
          "NetworkID": "e0bccd963bea8cbf9f8efe7a61dbd7d23e896820ef29e704389d6065f061ad21",
          "EndpointID": "8b581897bd6408f3836e1f60bc5bdfefe66695c0e2fd8ee127815f279a82f0d8",
          "Gateway": "172.17.0.1",
          "IPAddress": "172.17.0.2",
          "IPPrefixLen": 16,
          "IPv6Gateway": "",
          "GlobalIPv6Address": "",
          "GlobalIPv6PrefixLen": 0,
          "MacAddress": "02:42:ac:11:00:02",
          "DriverOpts": null
        }
      }
    },
    "Mounts": []
  }
 }
]
```

#### api: /ContainerOne
type:GET
功能：用于获取某个容器的信息
参数：

| 参数名 | 参数类型 | 示例 |
| - | :- | - |
| id | string |6fb453556811708c9ca162de467871cbb72cc787eb09e81f79cc0d713585d90c |

返回值：json
返回值示例：本版返回信息较多，后期可根据去求返回数据
>```
{
  "Id": "6fb453556811708c9ca162de467871cbb72cc787eb09e81f79cc0d713585d90c",
  "Created": "2018-07-30T08:42:32.704881503Z",
  "Path": "/bin/sh",
  "Args": [],
  "State": {
    "Status": "created",
    "Running": false,
    "Paused": false,
    "Restarting": false,
    "OOMKilled": false,
    "Dead": false,
    "Pid": 0,
    "ExitCode": 128,
    "Error": "driver failed programming external connectivity on endpoint tt2 (9b84c62a050af2477dfa2ffc071ffa7838e87d4afc15f301edba08ab01a8e073): Error starting userland proxy: listen tcp 0.0.0.0:8080: bind: address already in use",
    "StartedAt": "0001-01-01T00:00:00Z",
    "FinishedAt": "0001-01-01T00:00:00Z"
  },
  "Image": "sha256:11cd0b38bc3ceb958ffb2f9bd70be3fb317ce7d255c8a4c3f4af30e298aa1aab",
  "ResolvConfPath": "/var/lib/docker/containers/6fb453556811708c9ca162de467871cbb72cc787eb09e81f79cc0d713585d90c/resolv.conf",
  "HostnamePath": "",
  "HostsPath": "/var/lib/docker/containers/6fb453556811708c9ca162de467871cbb72cc787eb09e81f79cc0d713585d90c/hosts",
  "LogPath": "",
  "Name": "/tt2",
  "RestartCount": 0,
  "Driver": "overlay2",
  "Platform": "linux",
  "MountLabel": "",
  "ProcessLabel": "",
  "AppArmorProfile": "",
  "ExecIDs": null,
  "HostConfig": {
    "Binds": null,
    "ContainerIDFile": "",
    "LogConfig": {
      "Type": "json-file",
      "Config": {}
    },
    "NetworkMode": "default",
    "PortBindings": {
      "8080/tcp": [
        {
          "HostIp": "",
          "HostPort": "8080"
        }
      ]
    },
    "RestartPolicy": {
      "Name": "",
      "MaximumRetryCount": 0
    },
    "AutoRemove": false,
    "VolumeDriver": "",
    "VolumesFrom": null,
    "CapAdd": null,
    "CapDrop": null,
    "Dns": null,
    "DnsOptions": null,
    "DnsSearch": null,
    "ExtraHosts": null,
    "GroupAdd": null,
    "IpcMode": "shareable",
    "Cgroup": "",
    "Links": null,
    "OomScoreAdj": 0,
    "PidMode": "",
    "Privileged": false,
    "PublishAllPorts": false,
    "ReadonlyRootfs": false,
    "SecurityOpt": null,
    "UTSMode": "",
    "UsernsMode": "",
    "ShmSize": 67108864,
    "Runtime": "runc",
    "ConsoleSize": [
      0,
      0
    ],
    "Isolation": "",
    "CpuShares": 0,
    "Memory": 0,
    "NanoCpus": 0,
    "CgroupParent": "",
    "BlkioWeight": 0,
    "BlkioWeightDevice": null,
    "BlkioDeviceReadBps": null,
    "BlkioDeviceWriteBps": null,
    "BlkioDeviceReadIOps": null,
    "BlkioDeviceWriteIOps": null,
    "CpuPeriod": 0,
    "CpuQuota": 0,
    "CpuRealtimePeriod": 0,
    "CpuRealtimeRuntime": 0,
    "CpusetCpus": "",
    "CpusetMems": "",
    "Devices": null,
    "DeviceCgroupRules": null,
    "DiskQuota": 0,
    "KernelMemory": 0,
    "MemoryReservation": 0,
    "MemorySwap": 0,
    "MemorySwappiness": null,
    "OomKillDisable": false,
    "PidsLimit": 0,
    "Ulimits": null,
    "CpuCount": 0,
    "CpuPercent": 0,
    "IOMaximumIOps": 0,
    "IOMaximumBandwidth": 0
  },
  "GraphDriver": {
    "Data": {
      "LowerDir": "/var/lib/docker/overlay2/6cdd71bb1c988b4478455ac6fab4e54d9bdb852a9029be91f4aadb17fc3e5b86-init/diff:/var/lib/docker/overlay2/c21098a82e884844545bd7ffe79708abe3d5d90a51dcb85bd5bb01667dd512e8/diff",
      "MergedDir": "/var/lib/docker/overlay2/6cdd71bb1c988b4478455ac6fab4e54d9bdb852a9029be91f4aadb17fc3e5b86/merged",
      "UpperDir": "/var/lib/docker/overlay2/6cdd71bb1c988b4478455ac6fab4e54d9bdb852a9029be91f4aadb17fc3e5b86/diff",
      "WorkDir": "/var/lib/docker/overlay2/6cdd71bb1c988b4478455ac6fab4e54d9bdb852a9029be91f4aadb17fc3e5b86/work"
    },
    "Name": "overlay2"
  },
  "Mounts": [],
  "Config": {
    "Hostname": "6fb453556811",
    "Domainname": "",
    "User": "",
    "AttachStdin": true,
    "AttachStdout": false,
    "AttachStderr": false,
    "ExposedPorts": {
      "8080/tcp": {}
    },
    "Tty": false,
    "OpenStdin": true,
    "StdinOnce": false,
    "Env": [
      "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
    ],
    "Cmd": [
      "/bin/sh"
    ],
    "ArgsEscaped": true,
    "Image": "alpine",
    "Volumes": null,
    "WorkingDir": "",
    "Entrypoint": null,
    "OnBuild": null,
    "Labels": {}
  },
  "NetworkSettings": {
    "Bridge": "",
    "SandboxID": "dbec3854a5cb2611f252360624c29f65a8b9260e417921602f5de1704928c875",
    "HairpinMode": false,
    "LinkLocalIPv6Address": "",
    "LinkLocalIPv6PrefixLen": 0,
    "Ports": {},
    "SandboxKey": "/var/run/docker/netns/dbec3854a5cb",
    "SecondaryIPAddresses": null,
    "SecondaryIPv6Addresses": null,
    "EndpointID": "",
    "Gateway": "",
    "GlobalIPv6Address": "",
    "GlobalIPv6PrefixLen": 0,
    "IPAddress": "",
    "IPPrefixLen": 0,
    "IPv6Gateway": "",
    "MacAddress": "",
    "Networks": {
      "bridge": {
        "IPAMConfig": null,
        "Links": null,
        "Aliases": null,
        "NetworkID": "e0bccd963bea8cbf9f8efe7a61dbd7d23e896820ef29e704389d6065f061ad21",
        "EndpointID": "",
        "Gateway": "",
        "IPAddress": "",
        "IPPrefixLen": 0,
        "IPv6Gateway": "",
        "GlobalIPv6Address": "",
        "GlobalIPv6PrefixLen": 0,
        "MacAddress": "",
        "DriverOpts": null
      }
    }
  }
}
```

#### api: /ContainerStart
type:GET
功能：用于启动某个容器的
参数：

| 参数名 | 参数类型 | 示例 |
| - | :- | - |
| cid | string |725a6ee9a490cbdf02249ee0fbdd92bc2c3c29fa04b59c8631419d195eda97d2 |

返回值：string
返回值示例：“OK” （正确启动的返回值）


#### api: /CreateContainer
type:GET
功能：创建一个容器
参数：

| 参数名 | 参数类型 | 示例 |
| - | :- | - |
| name | string | Ctest |

返回值：string
返回值示例：“buildContainer” （正确创建的返回值）


#### api: /ImageList
type:GET
功能：获取所有镜像信息
参数：无

返回值：json
返回值示例：“buildContainer” （正确创建的返回值）
>```
{
  "Code": 0,
  "Msg": "",
  "Count": 6,
  "Data": [
    {
      "Containers": -1,
      "Created": 1532942551,
      "ID": "sha256:182e2fdb0e7d5923391ccaee47e276cc8278030c5c00c1d58272dfedc076465c",
      "Labels": "",
      "RepoDigests": null,
      "RepoTags": "ttst:dockertest",
      "Size": 28501298
    },
    {
      "Containers": -1,
      "Created": 1532452911,
      "ID": "sha256:c82521676580c4850bb8f0d72e47390a50d60c8ffe44d623ce57be521bca9869",
      "Labels": "NGINX Docker Maintainers <docker-maint@nginx.com>",
      "RepoDigests": [
        "nginx@sha256:d85914d547a6c92faa39ce7058bd7529baacab7e0cd4255442b04577c4d1f424"
      ],
      "RepoTags": "nginx:latest",
      "Size": 108975101
    },
    {
      "Containers": -1,
      "Created": 1532072033,
      "ID": "sha256:68872d9f77f8df9a66a08621372e594729d561fcc356d6979f62401a249a132e",
      "Labels": "",
      "RepoDigests": [
        "registry.cn-hangzhou.aliyuncs.com/alex-docker/alpine-go@sha256:70d7935a539ab884e1d4805bd5e11c8d61a2b274cea1c754772848d178c0c4ad"
      ],
      "RepoTags": "registry.cn-hangzhou.aliyuncs.com/alex-docker/alpine-go:v1.0",
      "Size": 12642032
    },
    {
      "Containers": -1,
      "Created": 1531252190,
      "ID": "sha256:caab7ec026902e98bf48505fa1e2c763cb9274dbf9268af36a78ff6de39a522a",
      "Labels": "Craig Citro <craigcitro@google.com>",
      "RepoDigests": [
        "tensorflow/tensorflow@sha256:92ad7f5da1f0e7c2c7b714b77b12424ae3d7971510d8ff8673b8b0695c3fd1c9"
      ],
      "RepoTags": "tensorflow/tensorflow:latest",
      "Size": 1251955358
    },
    {
      "Containers": -1,
      "Created": 1530886446,
      "ID": "sha256:11cd0b38bc3ceb958ffb2f9bd70be3fb317ce7d255c8a4c3f4af30e298aa1aab",
      "Labels": "",
      "RepoDigests": [
        "alpine@sha256:7043076348bf5040220df6ad703798fd8593a0918d06d3ce30c6c93be117e430"
      ],
      "RepoTags": "alpine:latest",
      "Size": 4413370
    },
    {
      "Containers": -1,
      "Created": 1528150774,
      "ID": "sha256:49f7960eb7e4cb46f1a02c1f8174c6fac07ebf1eb6d8deffbcb5c695f1c9edd5",
      "Labels": "",
      "RepoDigests": [
        "centos@sha256:b67d21dfe609ddacf404589e04631d90a342921e81c40aeaf3391f6717fa5322"
      ],
      "RepoTags": "centos:latest",
      "Size": 199678471
    }
  ]
}
```


#### api: /ImageOne
type:GET
功能：创建一个容器
参数：

| 参数名 | 参数类型 | 示例 |
| - | :- | - |
| image_id | string | sha256:c82521676580c4850bb8f0d72e47390a50d60c8ffe44d623ce57be521bca9869 |

返回值：json
返回值示例：本版返回信息较多，后期可根据去求返回数据
>```
{
  "Id": "sha256:c82521676580c4850bb8f0d72e47390a50d60c8ffe44d623ce57be521bca9869",
  "RepoTags": [
    "nginx:latest"
  ],
  "RepoDigests": [
    "nginx@sha256:d85914d547a6c92faa39ce7058bd7529baacab7e0cd4255442b04577c4d1f424"
  ],
  "Parent": "",
  "Comment": "",
  "Created": "2018-07-24T17:21:51.548456912Z",
  "Container": "895e85f09f69727097e9c1783362736d7ee9b5b32f9eae1e5e32f2e1002abf14",
  "ContainerConfig": {
    "Hostname": "895e85f09f69",
    "Domainname": "",
    "User": "",
    "AttachStdin": false,
    "AttachStdout": false,
    "AttachStderr": false,
    "ExposedPorts": {
      "80/tcp": {}
    },
    "Tty": false,
    "OpenStdin": false,
    "StdinOnce": false,
    "Env": [
      "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
      "NGINX_VERSION=1.15.2-1~stretch",
      "NJS_VERSION=1.15.2.0.2.2-1~stretch"
    ],
    "Cmd": [
      "/bin/sh",
      "-c",
      "#(nop) ",
      "CMD [\"nginx\" \"-g\" \"daemon off;\"]"
    ],
    "ArgsEscaped": true,
    "Image": "sha256:f1eea4ec6bee804c269e8443513d7afe26adb1615518ad56d014973fd5faa5f3",
    "Volumes": null,
    "WorkingDir": "",
    "Entrypoint": null,
    "OnBuild": [],
    "Labels": {
      "maintainer": "NGINX Docker Maintainers <docker-maint@nginx.com>"
    },
    "StopSignal": "SIGTERM"
  },
  "DockerVersion": "17.06.2-ce",
  "Author": "",
  "Config": {
    "Hostname": "",
    "Domainname": "",
    "User": "",
    "AttachStdin": false,
    "AttachStdout": false,
    "AttachStderr": false,
    "ExposedPorts": {
      "80/tcp": {}
    },
    "Tty": false,
    "OpenStdin": false,
    "StdinOnce": false,
    "Env": [
      "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
      "NGINX_VERSION=1.15.2-1~stretch",
      "NJS_VERSION=1.15.2.0.2.2-1~stretch"
    ],
    "Cmd": [
      "nginx",
      "-g",
      "daemon off;"
    ],
    "ArgsEscaped": true,
    "Image": "sha256:f1eea4ec6bee804c269e8443513d7afe26adb1615518ad56d014973fd5faa5f3",
    "Volumes": null,
    "WorkingDir": "",
    "Entrypoint": null,
    "OnBuild": [],
    "Labels": {
      "maintainer": "NGINX Docker Maintainers <docker-maint@nginx.com>"
    },
    "StopSignal": "SIGTERM"
  },
  "Architecture": "amd64",
  "Os": "linux",
  "Size": 108975101,
  "VirtualSize": 108975101,
  "GraphDriver": {
    "Data": {
      "LowerDir": "/var/lib/docker/overlay2/c520eac3cb1197b2c73e38fa08747583657cd45d177e26b6084886b1b7a5892a/diff:/var/lib/docker/overlay2/0532272d83bcf395c5b3f3a54c107ef82a9239081f767b1c002b97cc4388015c/diff",
      "MergedDir": "/var/lib/docker/overlay2/52422a38506423e561dacac6317cb9c6dee3a5cc4d76ad3fbc38216f2d6130c2/merged",
      "UpperDir": "/var/lib/docker/overlay2/52422a38506423e561dacac6317cb9c6dee3a5cc4d76ad3fbc38216f2d6130c2/diff",
      "WorkDir": "/var/lib/docker/overlay2/52422a38506423e561dacac6317cb9c6dee3a5cc4d76ad3fbc38216f2d6130c2/work"
    },
    "Name": "overlay2"
  },
  "RootFS": {
    "Type": "layers",
    "Layers": [
      "sha256:cdb3f9544e4c61d45da1ea44f7d92386639a052c620d1550376f22f5b46981af",
      "sha256:a8c4aeeaa0451a16218376ce6ec0e55094128baeb0dbe122f1b25c3fa81a5a5b",
      "sha256:08d25fa0442e3ea585b87bc6e9d41a1aa51624c83aec7fbafc1636f22eecf36f"
    ]
  },
  "Metadata": {
    "LastTagTime": "0001-01-01T00:00:00Z"
  }
}
```


#### api: /ImageSave
type:GET
功能：预留接口
参数：无

返回值：string
返回值示例：“ImageSave”
