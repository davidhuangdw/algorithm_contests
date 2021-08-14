package google_kickstart.`2019`.G

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, H) = input_ints()
        val A = input_longs()
        val B = input_longs()

        fun testset2(): Long {
            fun build(va: List<Long>, vb: List<Long>): MutableList<Pair<Long, Long>> {
                val scores = mutableListOf<Pair<Long, Long>>()
                fun dfs(i: Int, sa: Long, sb: Long) {
                    if (i == va.size) {
                        scores.add(sa to sb)
                    } else {
                        dfs(i + 1, sa + va[i], sb + vb[i])
                        dfs(i + 1, sa + va[i], sb)
                        dfs(i + 1, sa, sb + vb[i])
                    }
                }
                dfs(0, 0, 0)
                scores.sortBy { it.first }
                return scores
            }

            val hf = (N + 1) / 2
            val sl = build(A.subList(0, hf), B.subList(0, hf))
            val sr = build(A.subList(hf, A.size), B.subList(hf, B.size))
            if (sr.isEmpty()) {
                sr.add(0L to 0L)
            }

            val id = hashMapOf<Long, Int>()
            val pbs = sr.map { it.second }.toHashSet().sorted()
            for ((i, v) in pbs.withIndex())
                id[v] = i
            fun lowerId(x: Long): Int {
                var (l, r) = 0 to pbs.size - 1
                while (l <= r) {
                    val md = (l + r) / 2
                    if (pbs[md] <= x) l = md + 1
                    else r = md - 1
                }
                return r
            }

            val tree = MutableList<Long>(pbs.size + 1) { 0L }
            fun query(x: Long): Long {
                var i = lowerId(x) + 1
                var s = 0L
                while (i > 0) {
                    s += tree[i]
                    i -= i and -i
                }
                return s
            }

            fun add(x: Long) {   // add 1
                var i = id[x]!! + 1
                while (i < tree.size) {
                    tree[i] += 1L
                    i += i and -i
                }
            }

            var j = sr.size - 1
            var cnt = 0L
            var total = 0
            for ((ha, hb) in sl) {
                while (j >= 0 && sr[j].first >= H - ha) {
                    add(sr[j].second)
                    total += 1
                    j -= 1
                }
                cnt += total - query(H - hb - 1)
            }
            return cnt
        }

        fun testset1(): Long {
            var cnt = 0L
            fun dfs(i: Int, sa: Long, sb: Long) {
                if (i == N) {
                    if (sa >= H && sb >= H) cnt += 1
                } else {
                    dfs(i + 1, sa + A[i], sb + B[i])
                    dfs(i + 1, sa + A[i], sb)
                    dfs(i + 1, sa, sb + B[i])
                }
            }
            dfs(0, 0, 0)
            return cnt
        }

        println("Case #${it}: ${testset2()}")
    }
}
