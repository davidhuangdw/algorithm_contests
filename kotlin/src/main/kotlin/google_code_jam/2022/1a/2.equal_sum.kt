package google_code_jam.`2022`.`1a`

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    val K = 29
    val N = 100 // bug: forget to revert 100 after local test
    val A = (0L until N).toMutableList()
    for (i in 0 until N) {
        if (i <= K) {
            A[i] = 1L shl i
        } else {
            A[i] = A[i - 1] + 1   // bug: should be uniq
        }
    }
    val asum = A.sum()
    val line = A.joinToString(" ")
    (1..T).forEach {
        input_ints()
        println(line)
        System.out.flush()

        val B = input_longs()
        val C = (A.subList(minOf(N, K + 1), N) + B).sortedDescending()
        var remain = (asum + B.sum()) / 2
        val cur = mutableListOf<Long>()
        for (x in C) {
            if (x <= remain) {
                cur.add(x)
                remain -= x
            }
        }

        for (i in 0..K) {
            if (remain and (1L shl i) != 0L) {
                cur.add(A[i])
                remain -= A[i]
            }
        }
        if (remain != 0L) {
            throw RuntimeException("failed")
        }
        println(cur.joinToString(" "))
        System.out.flush()

//        println("Case #${it}: ")
    }
}
