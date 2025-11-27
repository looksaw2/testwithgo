#### 一些默认的配置需要修改的话可以改
import os
from pathlib import Path
import subprocess

##
rootPath : os.path = Path(__file__).parent.parent.parent
protoPath : os.path = "./api/proto"
orderprotoPath : os.path = Path(rootPath) / protoPath / "order.proto"
##

def generate_protoc() -> None:
    ### 开始生成order.proto
    commonDir: Path = Path(rootPath) / "internal" / "common"
    genProtoDir : Path = Path(commonDir) / "genproto"    
    if not genProtoDir.exists():
        genProtoDir.mkdir(parents=True , exist_ok=True)
    orderpbDir : Path = Path(genProtoDir) / "orderpb"
    if not orderpbDir.exists():
        orderpbDir.mkdir(parents=True , exist_ok=True)
        print(f"orderpb Dir has been created {orderpbDir}")
    else:
        print("orderpb Dir has already exist...........")
    ### 运行命令
    cmd = [
        "protoc", 
        f"--proto_path={Path(rootPath) / protoPath}",
        f"--go_out={genProtoDir} ",
        "--go_opt=module=github.com/looksaw2/gorder3/common/genproto/orderpb",
        f"--go-grpc_out={genProtoDir} ",
        "--go-grpc_opt=module=github.com/looksaw2/gorder3/common/genproto/orderpb",
        "order.proto"
    ]
    result = subprocess.run(cmd,capture_output=True,text=True)
    if result.returncode != 0:
        print(f"run the command error is {result.stderr}")
        return 
    print("执行命令成功................")
    


if __name__ == "__main__":
    generate_protoc()