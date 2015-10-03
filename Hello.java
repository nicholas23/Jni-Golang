public class Hello{
	static{
		System.loadLibrary("Hello");
	}
	public static native long add(long x ,long y);

	public static void main(String args[]){
		System.out.println("Hello");
		System.out.println(add(123,321));
	}
}
