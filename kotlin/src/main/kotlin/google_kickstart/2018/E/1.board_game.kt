fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N) = input_ints()
        val A = input_ints()
        val B = input_ints()
        val K = 3

        val pre = mutableListOf<Int>()
        fun gen(a: List<Int>, all: MutableList<List<Int>>) {
            val n = a.size
            if (n == K) {
                pre.add(a.sum())
                all.add(pre.toList())
                pre.removeLast()
            }
            for (i in 0..n - 3)
                for (j in i + 1..n - 2)
                    for (k in j + 1..n - 1) {
                        val chosen = listOf(i, j, k)
                        pre.add(chosen.map { a[it] }.sum())
                        val rem = (0 until n).filter { it !in chosen }.map { a[it] }
                        gen(rem, all)
                        pre.removeLast()
                    }
        }

        val pa = mutableListOf<List<Int>>()
        gen(A, pa)
        val pb = mutableListOf<List<Int>>()
        gen(B, pb)
        var max = 0.0
        for (ca in pa) {
            var total = 0
            var win = 0
            for (cb in pb) {
                total += 1
                var w = 0
                for ((x, y) in ca.zip(cb))
                    if (x > y)
                        w += 1
                if (w >= 2)
                    win += 1
            }
            max = maxOf(max, 1.0 * win / total)
        }
        // todo: testset2
        println("Case #${it}: $max")
    }
}
