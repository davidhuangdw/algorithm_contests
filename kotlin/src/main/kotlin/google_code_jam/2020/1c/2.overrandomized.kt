package google_code_jam.`2020`.`1c`

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    fun firstDig(x: Long): Int {
        var v = x
        while (v >= 10) {
            v /= 10
        }
        return v.toInt()
    }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (U) = input_ints()
        val all = hashSetOf<Char>()
        val freq = hashMapOf<Char, Int>()
        val cap = List(10) { hashSetOf<Char>() }
        for (_t in 1..10_000) {
            val row = split_input()
            val q = row[0].toLong()
            val r = row[1]
            for (ch in r)
                all.add(ch)
            val dig = r[0]
            if (q > 0) {
                cap[firstDig(q)].add(dig)
            }
            freq[dig] = freq.getOrDefault(dig, 0) + 1
        }

        val res = MutableList(10) { ' ' }
        for (ch in all)
            if (ch !in freq) {
                freq[ch] = 0
            }
        res[0] = freq.keys.minByOrNull { freq[it]!! }!!
        freq.remove(res[0])
        for (d in 1..9) {
            if (cap[d].size > 0) {
                res[d] = cap[d].maxByOrNull { freq.getOrDefault(it, -1) }!!
                freq.remove(res[d])
            } else {
                break
            }
        }
        var j = 0
        for ((i, ch) in res.withIndex())
            if (ch == ' ') {
                j = i
                break
            }

        for (ch in freq.keys.sortedByDescending { freq[it]!! }) {
            res[j++] = ch
        }
        println("Case #${it}: ${res.joinToString("")}")
    }
}
