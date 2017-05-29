# MiPlant

## Scan BLE devices
```
sudo timeout -s SIGINT 5s hcitool -i hci0 lescan | grep "Flower care" > scan.txt
```
