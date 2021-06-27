package google_kickstart.`2020`.A

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    data class Node(var c: Int = 0) {
        val nxt = mutableMapOf<Char, Node>()
        fun next(ch: Char): Node {
            if (!nxt.contains(ch)) nxt[ch] = Node()
            return nxt[ch]!!
        }
    }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, K) = input_ints()
        var res = 0

        val root = Node()
        repeat(N) {
            var node = root
            for (ch in readLine()!!)
                node = node.next(ch)
            node.c += 1
        }
        fun dfs(u: Node, h: Int): Int {
            var cnt = u.c
            for ((ch, v) in u.nxt) cnt += dfs(v, h + 1)
            res += cnt / K * h
            return cnt % K
        }
        dfs(root, 0)
        println("Case #${it}: $res")
    }
}
