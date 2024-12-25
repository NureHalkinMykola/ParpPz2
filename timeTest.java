import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

public class timeTest {

    public static void main(String[] args) {
        int size = 1000000000;
        int[] data = new int[size];
        for (int i = 0; i < size; i++) {
            data[i] = 1;
        }

        int threads = 5;
        ExecutorService executor = Executors.newFixedThreadPool(threads);
        long[] results = new long[threads];
        int chunkSize = size / threads;

        long startTime = System.nanoTime();
        for (int i = 0; i < threads; i++) {
            final int index = i;
            final int start = index * chunkSize;
            final int end = (index == threads - 1) ? size : start + chunkSize;

            executor.execute(() -> {
                long sum = 0;
                for (int j = start; j < end; j++) {
                    sum += data[j];
                }
                results[index] = sum;
            });
        }

        executor.shutdown();
        while (!executor.isTerminated()) {}

        long sum = 0;
        for (long s : results) {
            sum += s;
        }
        long endTime = System.nanoTime();

        System.out.println("Threads: " + threads);
        System.out.println("Sum: " + sum);
        System.out.println("Time: " + (endTime - startTime) / 1000000 + " ms");
    }
}
