fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    fun generate_list(line: List<Int>, n: Int): List<Int> {
        val vs = MutableList(n) { if (it < 2) line[it] else 0 }
        val (A, B, C, M) = line.subList(2, line.size)
        for (i in 2 until n)
            vs[i] = ((1L * vs[i - 1] * A % M + 1L * vs[i - 2] * B % M + C) % M).toInt()
        return vs
    }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, Q) = input_ints()
        val X = generate_list(input_ints(), N)
        val Y = generate_list(input_ints(), N)
        val Z = generate_list(input_ints(), Q)

        val L = (0 until N).map { minOf(X[it], Y[it]) + 1 }
        val R = (0 until N).map { maxOf(X[it], Y[it]) + 2 }
        var sum = 0L
        for(i in 0 until N){
            val l = minOf(X[i], Y[i])
            val r = maxOf(X[i], Y[i])
            sum += r - l + 1
        }
        val K = (1..Q).map { sum - Z[it - 1] to it }.sortedBy { it.first }

        var ret = 0L
        val points = hashMapOf<Int, Int>()
        for (l in L)
            points[l] = points.getOrDefault(l, 0) + 1
        for (r in R)
            points[r] = points.getOrDefault(r, 0) - 1
        var c = 0L
        var sc = 0L
        var l = -1
        var i = 0
        for (p in points.keys.sorted()) {
            val next = sc + (p - l) * c
            while (i < Q && K[i].first <= next) {
                val (k, idx) = K[i]
                if(c > 0){
                    val v = (k - sc + c - 1) / c + l - 1
                    ret += v * idx
                }
                i += 1
            }
            sc = next
            l = p
            c += points[p]!!
        }
        println("Case #${it}: $ret")
    }
}
