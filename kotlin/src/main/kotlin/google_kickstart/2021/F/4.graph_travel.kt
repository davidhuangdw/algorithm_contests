fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, M, K) = input_ints()
        val nodes = (0 until N).map { input_ints() }
        val adj = MutableList(N) { 0 }
        repeat(M) {
            val (i, j) = input_ints()
            adj[i] += 1 shl j
            adj[j] += 1 shl i
        }
        val U = 1 shl N

        val pow = hashMapOf<Int, Int>()
        for (i in 0 until N)
            pow[1 shl i] = i
        val points = MutableList(U) { 0L }
        val level = MutableList(N + 1) { hashSetOf<Int>() }
        val lev = MutableList(U) { 0 }
        level[0].add(0)
        lev[0] = 0

        for (b in 1 until U) {
            val lb = b and -b
            val st = b - lb
            lev[b] = lev[st] + 1
            level[lev[b]].add(b)
            points[b] = points[st] + nodes[pow[lb]!!][2]
        }

        val count = MutableList(U) { 0L }
        count[0] = 1
        var sum = 0L
        for (i in 0..N) {
            for (b in level[i]) {
                if (count[b] == 0L) continue
                val p = points[b]
                if (p == K.toLong()) {
                    sum += count[b]
                }

                if (b == 0) {
                    for (j in 0 until N) {
                        val t = 1 shl j
                        count[t] = 1
                    }
                } else {
                    for (j in 0 until N) {
                        if (b and (1 shl j) == 0 && p in nodes[j][0]..nodes[j][1] && adj[j] and b > 0) {
                            val t = b + (1 shl j)
                            count[t] += count[b]
                        }
                    }
                }
            }
        }

        println("Case #${it}: $sum")
    }
}
