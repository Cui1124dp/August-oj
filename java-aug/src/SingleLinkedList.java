import java.io.BufferedInputStream;
import java.util.Scanner;

/**
 * Created with IntelliJ IDEA.
 *
 * @Author: cdx1124dp
 * @Date: 2022/08/09/17:23
 * @Description:
 */
public class SingleLinkedList {

    private static int idx = 0;
    private static int head = -1;
    private static int[] val = new int[100010];
    private static int[] next = new int[100010];

    public static void insertHead(int x){
        val[idx] = x;
        next[idx] = head;
        head = idx;
        idx++;
    }

    public static void delete(int k){
        next[k] = next[next[k]];
    }

    public static void insert(int k,int x){
        val[idx] = x;
        next[idx] = next[k];
        next[k] = idx;
        idx++;
    }

    public static void printList(){

        int a = head;
        while(a!=-1){
            System.out.print(val[a]+" ");
            a = next[a];
        }

    }

    public static void main(String[] args) {

        Scanner scanner = new Scanner(new BufferedInputStream(System.in));
        int m = scanner.nextInt();
        while(m>=0){
            String[] s = scanner.nextLine().split(" ");
            if (s[0].equals("H")) insertHead(Integer.parseInt(s[1]));
            else if (s[0].equals("D")){
                if (Integer.parseInt(s[1])==0){
                    head = next[head];
                }else {
                    delete(Integer.parseInt(s[1])-1);
                }
            } else if (s[0].equals("I")) insert(Integer.parseInt(s[1])-1,Integer.parseInt(s[2]));
            m--;
        }
        printList();
    }

}
