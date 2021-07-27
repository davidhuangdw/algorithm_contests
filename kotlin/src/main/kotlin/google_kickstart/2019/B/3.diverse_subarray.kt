package google_kickstart.`2019`.B

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, S) = input_ints()
        val A = input_ints()
        var max = 0
        for (i in 0 until N) {
            var sum = 0
            val count = hashMapOf<Int, Int>()
            for (j in i until N) {
                val a = A[j]
                count[a] = count.getOrDefault(a, 0) + 1
                if (count[a]!! <= S) sum += 1
                else if (count[a]!! == S + 1) sum -= S
                max = maxOf(max, sum)
            }
        }
        println("Case #${it}: $max")
    }
    // todo: testset2
}
