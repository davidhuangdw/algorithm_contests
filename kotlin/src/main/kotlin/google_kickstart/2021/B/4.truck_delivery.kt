package google_kickstart.`2021`.B
fun main(){

    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().map{it.toInt()}
    fun input_longs() = split_input().map{it.toLong()}

    data class Edge(val to: Int, val load: Int, val toll: Long)

    fun gcd(aa: Long, bb: Long): Long{
        var (a, b) = aa to bb
        while(b != 0L){
            a = b.also{b = a % b}
        }
        return a
    }

    class SegmentG(val ML: Int) {
        val tree = LongArray(ML * 4) { 0L }

        fun _get_mch(x: Int, lx: Int, rx: Int) = listOf(lx + rx shr 1, x shl 1, x shl 1 or 1)

        fun set(pos: Int, v: Long) {
            fun _set(x: Int, lx: Int, rx: Int) {
                if (pos !in lx..rx) return
                val (m, lch, rch) = _get_mch(x, lx, rx)
                if(lx == rx){
                    tree[x] = v
                    return
                }
                if (pos <= m) {
                    _set(lch, lx, m)
                } else {
                    _set(rch, m + 1, rx)
                }
                tree[x] = gcd(tree[lch], tree[rch])
            }
            _set(1, 1, ML)
        }

        fun query(l: Int, r: Int): Long{
            fun _query(x: Int, lx: Int, rx: Int): Long{
                if(rx < l || r < lx) return 0L
                if(l <= lx && rx <= r) return tree[x]
                val (m, lch, rch) = _get_mch(x, lx, rx)
                return gcd(_query(lch, lx, m), _query(rch, m+1, rx))
            }
            return _query(1, 1, ML)
        }
    }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, Q) = input_ints()
        val Adj = hashMapOf<Int, MutableList<Edge>>()
        val Queries = hashMapOf<Int, MutableList<Pair<Int, Int>>>()
        var max_load = 0
        repeat(N-1){
            var (xx, yy, lload, toll) = input_longs()
            val (x, y, load) = listOf(xx, yy, lload).map{it.toInt()}
            Adj.getOrPut(x){mutableListOf()}.add(Edge(y, load, toll))
            Adj.getOrPut(y){mutableListOf()}.add(Edge(x, load, toll))
            max_load = maxOf(max_load, load)
        }
        (0 until Q).forEach{ i ->
            val (x, w) = input_ints()
            Queries.getOrPut(x){mutableListOf()}.add(i to w)
        }

        val seg = SegmentG(max_load)
        val ans = LongArray(Q)
        fun dfs(v: Int, parent: Int){
            for((i, w) in Queries.getOrElse(v){listOf()}){
                ans[i] = seg.query(1, w)
            }
            for((to, load, toll) in Adj.getOrElse(v){listOf()}){
                if(to == parent) continue
                seg.set(load, toll)
                dfs(to, v)
                seg.set(load, 0L)
            }
        }
        dfs(1, -1)
        println("Case #${it}: ${ans.joinToString(" ")}")
    }
}
