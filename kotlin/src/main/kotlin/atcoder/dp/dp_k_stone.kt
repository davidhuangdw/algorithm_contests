package atcoder.dp

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val (N, K) = input_ints()
    val A = input_ints()
    val win = BooleanArray(K + 1)
    for (k in 0..K) {
        win[k] = false
        for (a in A) {
            if (a > k) break
            if (!win[k-a]) {
                win[k] = true
                break
            }
        }
    }
    println(if(win[K]) "First" else "Second")
}
