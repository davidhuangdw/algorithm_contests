package google_kickstart.`2020`.C

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    class FenwickTree(arr: List<Long>) {
        val N = arr.size
        val tree = LongArray(N + 1)

        init {
            val pre = arr.scan(0, Long::plus)
            for (i in 1..N)
                tree[i] = pre[i] - pre[i and i - 1]
        }

        fun query(i: Int): Long {
            var (res, j) = 0L to i
            while (j > 0) {
                res += tree[j]
                j -= j and -j
            }
            return res
        }

        fun add(i: Int, v: Long) {
            var j = i
            while (j <= N) {
                tree[j] += v
                j += j and -j
            }
        }

        fun update(i: Int, v: Long) = add(i, v - (query(i) - query(i - 1)))

        fun debug() {
            println((1..N).map { query(it) })
        }
    }

    fun neg(i: Int) = if (i % 2 == 0) 1 else -1

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, Q) = input_ints()
        val A = input_longs()
        val fw1 = FenwickTree(A.mapIndexed { i, v -> neg(i + 1) * v })
        val fw2 = FenwickTree(A.mapIndexed { i, v -> neg(i + 1) * v * (i + 1) })

        var sum = 0L
        repeat(Q) {
            val (O, X, Y) = split_input()
            if (O == "U") {
                val i = X.toInt()
                val v = Y.toLong()
                fw1.update(i, neg(i) * v)
                fw2.update(i, neg(i) * v * i)
            } else {
                val l = X.toInt()
                val r = Y.toInt()
                val sub1 = fw1.query(r) - fw1.query(l - 1)
                val sub2 = fw2.query(r) - fw2.query(l - 1)
                val cur = neg(-l) * (sub2 + (1 - l) * sub1)
                sum += cur
            }
        }
        println("Case #${it}: $sum")
    }
}
