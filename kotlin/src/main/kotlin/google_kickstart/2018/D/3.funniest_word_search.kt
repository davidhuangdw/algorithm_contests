fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    fun gcd(a: Int, b: Int): Int = if (a == 0) b else gcd(b % a, a)

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (R, C, W) = input_ints()
        val A = (1..R).map { readLine()!! }
        val dict = (1..W).map { readLine()!! }
        fun testset1(): String {
            val ds = dict.map { it[0] }.toHashSet()
            val map = A.map { it.map { ch -> if (ch in ds) 1 else 0 } }
            var cnt = 0
            var mx = 0 to 1
            for (i in 0 until R) {
                val row = MutableList(C) { 0 }
                for (j in i until R) {
                    for (k in 0 until C)
                        row[k] += map[j][k]
                    val pre = row.scan(0, Int::plus)
                    val h = j - i + 1
                    for (l in 0 until C)
                        for (r in l + 1..C) {
                            val y = (pre[r] - pre[l]) * 4
                            val z = h + r - l
                            if (y * mx.second == z * mx.first) {
                                cnt += 1
                            } else if (y * mx.second > z * mx.first) {
                                val g = gcd(y, z)
                                mx = y / g to z / g
                                cnt = 1
                            }
                        }
                }
            }
            return "${mx.first}/${mx.second} $cnt"
        }

        fun testset2(): String { //todo: fix it - RE
            val count = hashMapOf<String, Int>()
            for (w in dict.toList()) {
                count[w] = count.getOrDefault(w, 0) + 1
                val rw = w.reversed()
                count[rw] = count.getOrDefault(rw, 0) + 1
            }
            val rows = List(R) { List(C) { MutableList(C) { 0 } } }
            val cols = List(C) { List(R) { MutableList(R) { 0 } } }
            var cnt = 0
            var mx = 0 to 1
            for (i in 0 until R)
                for (j in 0 until C) {
                    for (k in i until R) {
                        val w = (i..k).map { A[it][j] }.joinToString("")
                        val c = count.getOrDefault(w, 0)
                        if (c > 0) {
                            for (x in 0..i)
                                cols[j][x][k] += c * w.length
                        }
                    }
                    for (k in j until C) {
                        val w = (j..k).map { A[i][it] }.joinToString("")
                        val c = count.getOrDefault(w, 0)
                        if (c > 0) {
                            for (x in 0..i)
                                rows[i][x][k] += c * w.length
                        }
                    }
                }

            for (i in 0 until R) {
                val down = MutableList(C) { 0 }
                val right = List(C) { MutableList(C) { 0 } }
                for (j in i until R) {
                    for (k in 0 until C)
                        down[k] += cols[k][i][j]
                    val h = j - i + 1
                    for (l in 0 until C) {
                        var y = 0
                        for (r in l until C) {
                            right[l][r] += rows[j][l][r]
                            y += down[r] + right[l][r]
                            val z = h + r - l + 1
                            if (y * mx.second == z * mx.first) {
                                cnt += 1
                            } else if (y * mx.second > z * mx.first) {
                                val g = gcd(y, z)
                                mx = y / g to z / g
                                cnt = 1
                            }
                        }
                    }
                }
            }
            return "${mx.first}/${mx.second} $cnt"
        }
        println("Case #${it}: ${testset2()}")
    }
}
