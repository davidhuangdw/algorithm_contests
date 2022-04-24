package google_code_jam.`2020`.r2

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val MAX = 1000_000
    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (C, D) = input_ints()
        val X = listOf(0) + input_ints()
        val adj = List(C) { mutableListOf<Int>() }
        val edge = hashMapOf<Pair<Int, Int>, Int>()
        for (i in 0 until D) {
            var (a, b) = input_ints()
            a--
            b--
            adj[a].add(b)
            adj[b].add(a)
            edge[a to b] = i
            edge[b to a] = i
        }

        val res = MutableList(D) { MAX }
        val latency = MutableList(C) { i -> if (i == 0) 0 else -1 }

        fun set1() {
            val xInd = hashMapOf<Int, MutableList<Int>>()
            for ((i, _x) in X.withIndex()) {
                val x = -_x
                if (x !in xInd) {
                    xInd[x] = mutableListOf()
                }
                xInd[x]!!.add(i)
            }
            var lat = 1
            for (x in xInd.keys.sorted()) {
                if (x == 0)
                    continue
                for (i in xInd[x]!!) {
                    latency[i] = lat
                    val j = adj[i].first { j -> latency[j] >= 0 }
                    res[edge[j to i]!!] = latency[i] - latency[j]
                }
                lat++
            }
        }

        fun set2() {
            val preToInd = hashMapOf<Int, MutableList<Int>>()
            val latToInd = hashMapOf<Int, MutableList<Int>>()
            for ((i, _x) in X.withIndex()) {
                if (_x > 0) {
                    val x = _x
                    if (x !in latToInd) latToInd[x] = mutableListOf()
                    latToInd[x]!!.add(i)
                } else if (_x < 0) {
                    val x = -_x
                    if (x !in preToInd) preToInd[x] = mutableListOf()
                    preToInd[x]!!.add(i)
                }
            }
            val lats = latToInd.keys.sorted()
            var li = 0
            val preCounts = preToInd.keys.sorted()
            var pre = 1
            var preLat = 0
            var pi = 0
            while (pi < preCounts.size) {
                var ind: List<Int>
                val lat: Int
                val c = preCounts[pi]
                if (pre != c) {
                    if (pre > c)
                        throw RuntimeException("not satisfy")
                    lat = lats[li]
                    ind = latToInd[lat]!!
                    li++
                } else {
                    lat = preLat + 1
                    ind = preToInd[c]!!
                    pi++
                }
                for (i in ind) {
                    latency[i] = lat
                    val j = adj[i].first { j -> latency[j] in 0 until lat }
                    res[edge[j to i]!!] = latency[i] - latency[j]
                }
                pre += ind.size
                preLat = lat
            }

            for (lat in lats.subList(li, lats.size)) {
                for (i in latToInd[lat]!!) {
                    latency[i] = lat // bug: missing
                    val j = adj[i].first { j -> latency[j] in 0 until lat }
                    res[edge[j to i]!!] = latency[i] - latency[j]
                }
            }
        }
        set2()

        println("Case #${it}: ${res.joinToString(" ")}")
    }
}
