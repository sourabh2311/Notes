## Build

```
$> cd avi-dev/buildconfig
avi-dev/buildconfig $> sudo python devinstall.py
avi-dev $> mkdir build
avi-dev $> cd build
avi-dev/build $> cmake ../
avi-dev/build $> make -j6
avi-dev/build $> make install_avi
avi-dev/build $> make controller_ova
```

Among these mainly `make -j6` takes time.

_Note: It might be possible that `make install_avi` fails, in that case, in avi-dev, run the script fix_go_build....sh_

## Deploy

Navigate to avi-dev/test/robot/new/lib/tools/

Run `python setup_env.py -n ~/my_tb.conf -o ~/avi-dev/build`

_Note: There came an error cURL something, solution to that was to update ovftools, just download the latest one and then install it_

## AWS Creation

Go to Infrastructure, Clouds, Create and then see the video

## Virtual Service Creation

Go to Applications, Virtual Services, Create Virtual Service and then see the video

## Misc

- At times it may happen that you made changes to changes to your go files, what you can do in that case is go to `~/avi-dev/build` and type `cmake ../` and then `make go_controller` which will then create a binary in `~/avi-dev/build/bin` named "go_controller" which can now be copied like shown in video

- To access my deployed, go to 10.52.0.73, username and password both are admin

- This is the only line to be written in terminal at avi/infra/ to get the dependencies straight if using a new package in go, and maybe this will cause the change in the file Gopkd.lock `bash dep.sh ensure`

- Run `pbtest` to run PR sanity check.

- If changing files on in `go` folder, then one can do:
  1. `service process-supervisor stop`
  2. `cd /opt/avi/bin/`
  3. `mv go-controller go-controller-orig`
  4. `mv /home/admin/go-controller /opt/avi/bin/`
  5. `service process-supervisor start`
  6. `service cloudconnectorgo status`
