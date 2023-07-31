package google_code_jam.`2019`.`1c`

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val win = hashMapOf('R' to 'P', 'P' to 'S', 'S' to 'R')
    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (A) = input_ints()
        val S = (1..A).map { split_input()[0] }.toHashSet().toList()
        var s = (S.indices).toList()
        val p = mutableListOf<Char>()
        var succ = true
        while (s.isNotEmpty()) {
            val i = p.size
            val ch = s.map { S[it][i % S[it].length] }.toHashSet()
            if (ch.size >= 3) {
                succ = false
                break
            }
            if (ch.size == 1) {
                p.add(win[ch.firstOrNull()!!]!!)
                break
            }
            var (a, b) = ch.toList()
            if (win[a] == b)
                a = b
            s = s.filter { S[it][i % S[it].length] == a }
            p.add(a)
        }
        val res = if (succ) p.joinToString("") else "IMPOSSIBLE"
        println("Case #${it}: $res")
    }
}
