import java.util.Arrays;
import java.util.Scanner;

public class SmartPhone {
	public static void main(String args[]) {
		int n;
		Scanner sc = new Scanner(System.in);
		n = sc.nextInt();
		//System.out.println(n);
		long[] customerBudget = new long[n];
		for (int i=0;i<n;i++) {
			//System.out.println(i);
			customerBudget[i] = sc.nextLong();
			//System.out.println(sc.nextLong());
			sc.skip("(\r\n|[\n\r\u2028\u2029\u0085])?");
		}
		int length = customerBudget.length;
		//System.out.println(length);
		long max = 0;
		Arrays.sort(customerBudget);
		for(long val : customerBudget) {
			//System.out.println(val);
			if( max < val*length) {
				max = val * length;
			}
			length -- ;
		}
		System.out.println(max);
	}
}
