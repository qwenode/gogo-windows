package memory

import (
	"bytes"
	"github.com/0xrawsec/golang-win32/win32"
	"github.com/0xrawsec/golang-win32/win32/kernel32"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
)

type ProcessInfo struct {
	Pid         uint32          //进程ID
	Path        string          //进程所在目录
	Handle      win32.HANDLE    //句柄
	BaseAddress win32.ULONGLONG //进程基址
	RegionSize  win32.ULONGLONG //进程分配空间大小
}

type BaseMemoryBlock struct {
	Handle      win32.HANDLE    //句柄
	BaseAddress win32.ULONGLONG //内存基址
	RegionSize  win32.ULONGLONG //内存空间大小
}

func OpenProcessByPid(pid uint32) (ProcessInfo, error) {
	process, err := kernel32.OpenProcess(win32.GENERIC_ALL, win32.BOOL(0), win32.DWORD(pid))
	if err != nil {
		return ProcessInfo{}, err
	}
	path, _ := kernel32.QueryFullProcessImageName(process)
	p := ProcessInfo{
		Pid:    pid,
		Path:   path,
		Handle: process,
	}
	return p, nil
}

// DecodeWindowString decodes MS-Windows (16LE) UTF files
func DecodeWindowString(b []byte) string {
	utf16 := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
	decoder := utf16.NewDecoder()
	reader := transform.NewReader(bytes.NewReader(b), unicode.BOMOverride(decoder))
	all, _ := io.ReadAll(reader)
	return string(all)
}
