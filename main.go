package main
//#cgo CFLAGS: -I/usr/lib/jvm/java-8-oracle/include
//#cgo CFLAGS: -I/usr/lib/jvm/java-8-oracle/include/linux
//#include <jni.h>
import "C"
import "fmt"

//export Java_Hello_add
func Java_Hello_add(env *C.JNIEnv, clazz C.jclass, x C.jlong, y C.jlong) C.jlong {
	fmt.Println("Hello from Go");
	return x + y
}

func main(){}
