package google_code_jam.`2018`.`1a`

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val (T) = input_ints()
    (1..T).forEach {
        val (R, C, H, V) = input_ints()
        val A = (1..R).map { split_input()[0] }
        val dinners = (H + 1) * (V + 1)
        val pieces = A.sumOf { it.sumOf { if (it == '@') 1L else 0L } }.toInt()
        fun set2(): Boolean {
            val pre = List(R + 1) { MutableList(C + 1) { 0 } }
            for (i in 1..R)
                for (j in 1..C) {
                    pre[i][j] = pre[i][j - 1] + pre[i - 1][j] - pre[i - 1][j - 1] + if (A[i - 1][j - 1] == '@') 1 else 0
                }
            if (pieces % dinners != 0) return false
            val unit = pieces / dinners
            val rs = MutableList(H + 2) { 0 }
            val cs = MutableList(V + 2) { 0 }
            for (k in 1..H + 1) {
                val t = k * unit * (V + 1)
                for (i in 1..R) {
                    val c = pre[i][C]
                    if (c == t) {
                        rs[k] = i
                        break
                    }
                    if (c > t) return false
                }
            }
            for (k in 1..V + 1) {
                val t = k * unit * (H + 1)
                for (i in 1..C) {
                    val c = pre[R][i]
                    if (c == t) {
                        cs[k] = i
                        break
                    }
                    if (c > t) return false
                }
            }
//            println(rs)
//            println(cs)
            for(i in 1..H+1)
                for(j in 1..V+1){
                    val (a, b) = rs[i-1]  to rs[i]
                    val (c, d) = cs[j-1]  to cs[j]
                    val cnt = pre[b][d] - pre[a][d] - pre[b][c] + pre[a][c]
                    if(cnt != unit)
                        return false
                }
            return true
        }

        val res = if (set2()) "POSSIBLE" else "IMPOSSIBLE"
        println("Case #${it}: $res")
    }
}
