package google_kickstart.`2021`.D

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val A = (0 until 3).map {
            val row = input_longs()
            if (it == 1)
                listOf(row[0], -1, row[1])
            else row
        }

        fun ok(seq: List<Long>) = seq[1] - seq[0] == seq[2] - seq[1]

        var cnt = 0
        for (i in listOf(0, 2)) {
            val row = A[i]
            if (ok(row)) cnt += 1

            val col = (0 until 3).map { A[it][i] }
            if (ok(col)) cnt += 1
        }

        val count = hashMapOf<Long, Int>()
        var max = 0
        for ((a, c) in listOf(A[0][0] to A[2][2], A[0][1] to A[2][1], A[0][2] to A[2][0], A[1][0] to A[1][2])) {
            if ((a + c) % 2 != 0L) continue
            val b = (a + c) / 2
            count[b] = count.getOrDefault(b, 0) + 1
            max = maxOf(max, count[b]!!)
        }

        println("Case #${it}: ${cnt + max}")
    }
}
