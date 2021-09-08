fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val M = 26
        val (N) = input_ints()
        val A = readLine()!!
        val B = readLine()!!
        val count = MutableList(26) { 0 }
        val exist = hashSetOf<List<Int>>()
        for (i in 0 until N) {
            count.replaceAll { 0 }
            for (j in i until N) {
                count[B[j] - 'A'] += 1
                exist.add(count.toList())
            }
        }

        var cnt = 0
        for (i in 0 until N) {
            count.replaceAll { 0 }
            for (j in i until N) {
                count[A[j] - 'A'] += 1
                if (count in exist)
                    cnt += 1
            }
        }
        println("Case #${it}: $cnt")
    }
}
