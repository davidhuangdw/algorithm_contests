fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        print("Case #${it}: ")
        val (N, M, A, B) = input_ints()
        fun solve() {
            val min = N + M - 1
            val r = A - min
            val l = B - min
            if (r >= 0 && l >= 0) {
                println("Possible")
                val row = MutableList(M) { 1 }
                for (i in 0 until N - 1) {
                    println(row.joinToString(" "))
                }
                row[0] += l
                row[M - 1] += r
                println(row.joinToString(" "))
            } else {
                println("Impossible")
            }
        }
        solve()
    }
}
