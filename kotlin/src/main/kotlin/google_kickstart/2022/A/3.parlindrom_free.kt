package google_kickstart.`2022`.A

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    fun isPar(x: List<Int>): Boolean {
        val n = x.size
        var (i, j) = 0 to n - 1
        while (i < j) {
            if (x[i] != x[j])
                return false
            i++
            j--
        }
        return true
    }

    val Pars = MutableList<HashSet<Int>>(7) { hashSetOf() }
    for (l in 5..6) {
        for (v in 0 until (1 shl l)) {
            var vs = MutableList(l) { k -> if ((v and (1 shl k)) == 0) 0 else 1 }
            if (isPar(vs)) Pars[l].add(v)
        }
    }
    fun isPar(l: Int, vs: Int) = Pars[l].contains(vs)

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N) = input_ints()
        val (S) = split_input()
        var pre = mutableListOf(0)  // bug: initial value use 0~31

        for ((i, ch) in S.withIndex()) {
            var xs = mutableListOf<Int>()
            when (ch) {
                '?' -> {
                    xs.add(0)
                    xs.add(1)
                }
                else -> {
                    xs.add(ch - '0')
                }
            }
            var cur = mutableListOf<Int>()
            for (v in pre) {
                for (x in xs) {
                    cur.add((v shl 1) + x)
                }
            }
            if (i >= 5) {
                var tmp = cur
                cur = mutableListOf()
                for (v in tmp) {
                    if (!isPar(6, v)) cur.add(v and 31)
                }
            }
            if (i >= 4) {
                var tmp = cur
                cur = mutableListOf()
                for (v in tmp) {
                    if (!isPar(5, v)) cur.add(v)
                }
            }
            pre = cur
            if (pre.isEmpty()) break
        }
        var res = if (pre.isEmpty()) "IMPOSSIBLE" else "POSSIBLE"

        println("Case #${it}: $res")
    }
}
