import java.io.BufferedInputStream;
import java.util.Scanner;

/**
 * Created with IntelliJ IDEA.
 *
 * @Author: cdx1124dp
 * @Date: 2022/08/08/2:39
 * @Description:
 */
public class NumberOfReversePairs {

    private static long COUNT = 0;

    public static void margeSort(int[] num,int l,int r){
        if (l>=r){
            return;
        }
        int mid = (l+r)>>1;
        margeSort(num,l,mid);
        margeSort(num,mid+1,r);
        int[] temp = new int[r-l+1];
        int a=0,b=l,c=mid+1;
        while(b<=mid&&c<=r){
            if (num[b]<=num[c]){
                temp[a++]=num[b++];
            }else {
                temp[a++]=num[c++];
                COUNT += (mid-b+1);
            }
        }
        while (b<=mid){
            temp[a++]=num[b++];
        }
        while (c<=r){
            temp[a++]=num[c++];
        }
        int j = 0;
        for (int i = l; i <=r ; i++) {
            num[i] = temp[j++];
        }
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(new BufferedInputStream(System.in));
        int n = scanner.nextInt();
        int[] num = new int[n];
        for (int i = 0;i<n;i++){
            num[i] = scanner.nextInt();
        }
        margeSort(num,0,n-1);
        System.out.println(COUNT);
    }

}
