package main

import (
	"context"
	"log"
	"tiny-dfs/gen-go/tdfs"
)

type DataNodeHandler struct {
	core *DataNodeCore
}

func NewDataNodeHandler(core *DataNodeCore) *DataNodeHandler {
	return &DataNodeHandler{
		core: core,
	}
}

func (d *DataNodeHandler) MakeReplica(ctx context.Context, target_addr string, file_path string) (_r *tdfs.Response, _err error) {
	d.core.MakeReplica(target_addr, file_path)
	return &tdfs.Response{Status: 200, Msg: "MakeReplica ok"}, nil
}

func (d *DataNodeHandler) ReceiveReplica(ctx context.Context, file_path string, file *tdfs.File) (_r *tdfs.Response, _err error) {
	err := d.core.Save(file_path, file.Data, file.Medatada)
	if err != nil {
		log.Println("创建副本时新建文件失败：", err)
		return &tdfs.Response{Status: 500, Msg: "ReceiveReplica Failed"}, err
	}
	return &tdfs.Response{Status: 200, Msg: "ReceiveReplica ok"}, nil
}

func (d *DataNodeHandler) Ping(ctx context.Context) (_r *tdfs.DNStat, _err error) {
	return d.core.GetStat(), nil
}

func (d *DataNodeHandler) Heartbeat(ctx context.Context) (_r *tdfs.Response, _err error) {
	//TODO implement me
	panic("implement me")
}

func (d *DataNodeHandler) Put(ctx context.Context, remote_file_path string, file_data []byte, metadata *tdfs.Metadata) (_r *tdfs.Response, _err error) {
	log.Println("Enter hanlder Put")

	err := d.core.Save(remote_file_path, file_data, metadata)
	if err != nil {
		log.Println(err)
	}
	return &tdfs.Response{Status: 200, Msg: "Put ok"}, nil
}

func (d *DataNodeHandler) Get(ctx context.Context, remote_file_path string) (_r *tdfs.Response, _err error) {
	file, err := d.core.Get(remote_file_path)
	if err != nil {
		return &tdfs.Response{Status: 400}, err
	}
	return &tdfs.Response{Status: 200, File: file}, nil
}

func (d *DataNodeHandler) Delete(ctx context.Context, remote_file_path string) (_r *tdfs.Response, _err error) {
	d.core.Delete(remote_file_path)
	return &tdfs.Response{Status: 200}, nil
}
