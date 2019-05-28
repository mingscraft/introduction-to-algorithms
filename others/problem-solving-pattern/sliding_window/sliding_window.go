package sliding_window

func MaxSubArraySum(arr []int, num int) int {
    maxSum := 0
    tmpSum := 0

    if len(arr) < num {
        return 0
    }

    for i := 0; i < num; i ++ {
       maxSum += arr[i]
    }

    tmpSum = maxSum

    for i := num; i < len(arr); i++ {
        tmpSum = tmpSum - arr[i-num] + arr[i]
        if tmpSum > maxSum {
            maxSum = tmpSum
        }
    }

    return maxSum
}
