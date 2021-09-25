fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, P) = input_ints()
        var cnt = 1L shl N
        val ps = hashSetOf<String>()
        fun exist_prefix(s: String): Boolean {
            for (i in 1..s.length)
                if (s.substring(0, i) in ps)
                    return true
            return false
        }
        for (p in (1..P).map { readLine()!! }.sortedBy { it.length }) {
            if (exist_prefix(p)) continue
            ps.add(p)
            cnt -= 1L shl (N - p.length)
        }

        println("Case #${it}: $cnt")
    }
}
