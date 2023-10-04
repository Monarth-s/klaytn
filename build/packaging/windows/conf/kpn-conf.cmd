REM Configuration file for the kpn

set NETWORK_ID=8217

set PORT=32323

set SERVER_TYPE="http"
set SYNCMODE="full"
set VERBOSITY=3
set MAXCONNECTIONS=200
:: set LDBCACHESIZE=10240

REM txpool options setting
set TXPOOL_EXEC_SLOTS_ALL=8192
set TXPOOL_NONEXEC_SLOTS_ALL=8192
set TXPOOL_EXEC_SLOTS_ACCOUNT=8192
set TXPOOL_NONEXEC_SLOTS_ACCOUNT=8192
set TXPOOL_LIFE_TIME="5m"

REM rpc options setting
set RPC_ENABLE=0 &:: if this is set, the following options will be used
set RPC_API="klay" &:: available apis: admin,debug,klay,eth,miner,net,personal,rpc,txpool,web3
set RPC_PORT=8551
set RPC_ADDR="0.0.0.0"
set RPC_CORSDOMAIN="*"
set RPC_VHOSTS="*"

REM ws options setting
set WS_ENABLE=0 &:: if this is set, the following options will be used
set WS_API="klay" &:: available apis: admin,debug,klay,eth,miner,net,personal,rpc,txpool,web3
set WS_ADDR="0.0.0.0"
set WS_PORT=8552
set WS_ORIGINS="*"

REM Setting 1 is to enable options, otherwise disabled.
set METRICS=1
set PROMETHEUS=1
set DB_NO_PARALLEL_WRITE=0
set MULTICHANNEL=1

REM Raw options e.g) "--txpool.nolocals"
set ADDITIONAL="--txpool.nolocals"

set VINI_HOME=%homepath%\.kpn

set DATA_DIR=%VINI_HOME%\data
