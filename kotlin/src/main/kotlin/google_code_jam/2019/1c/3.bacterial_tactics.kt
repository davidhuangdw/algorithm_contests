package google_code_jam.`2019`.`1c`

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

//    File("debug.txt").printWriter().use{debug -> }
    val (T) = input_ints()
    (1..T).forEach {
        val (R, C) = input_ints()
        val A = (1..R).map { split_input()[0] }

        fun enc(i: Int, j: Int) = i * C + j
        fun dec(x: Int) = x / C to x % C
        fun isMut(x: Int) = A[x / C][x % C] == '#'
        fun isEmpty(st: Int, x: Int) = st and (1 shl x) == 0
        fun mark(st: Int, i: Int, j: Int) = st + (1 shl enc(i, j))

        var init = 0
        for (x in (0 until R * C))
            if (isMut(x))
                init = init or (1 shl x)

        fun horizon(st: Int, x: Int): Int {
            var y = st
            val (si, sj) = dec(x)
            var i = si
            var j = sj
            while (j >= 0 && isEmpty(st, enc(i, j))) {
                y = mark(y, i, j)
                j--
            }
            if (j >= 0 && isMut(enc(i, j))) return -1
            j = sj + 1
            while (j < C && isEmpty(st, enc(i, j))) {
                y = mark(y, i, j)
                j++
            }
            if (j < C && isMut(enc(i, j))) return -1
            if (y < st)
                throw RuntimeException("failed: $st $y")
            return y
        }

        fun vertical(st: Int, x: Int): Int {
            var y = st  // bug: not y=x
            val (si, sj) = dec(x)
            var i = si
            val j = sj
            while (i >= 0 && isEmpty(st, enc(i, j))) {
                y = mark(y, i, j)
                i--
            }
            if (i >= 0 && isMut(enc(i, j))) return -1
            i = si + 1
            while (i < R && isEmpty(st, enc(i, j))) {
                y = mark(y, i, j)
                i++
            }
            if (i < R && isMut(enc(i, j))) return -1
            if (y < st)
                throw RuntimeException("failed: $st $y")
            return y
        }

        val done = hashMapOf<Int, Int>()
        fun dfs(st: Int): Int {
            if (st !in done) {
                var cnt = 0
                for (x in 0 until R * C)
                    if (isEmpty(st, x)) {
                        var y = horizon(st, x)
                        if (y > 0 && dfs(y) == 0)
                            cnt++

                        y = vertical(st, x)
                        if (y > 0 && dfs(y) == 0)
                            cnt++
                    }
                done[st] = cnt
            }
            return done[st]!!
        }

        val res = dfs(init)
        // todo: set2
        println("Case #${it}: $res")
    }
}
