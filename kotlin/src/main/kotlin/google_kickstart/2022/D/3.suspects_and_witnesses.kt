package google_kickstart.`2022`.D
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    var (N, M, K) = listOf(0, 0, 0)

    val (T, ) = input_ints()
    (1..T).forEach {
        val (N, M, K) = input_ints()
        val edge = hashMapOf<Int, MutableList<Int>>()
        (1..M).forEach {
            val (a, b) = input_ints()
            if (b !in edge){
                edge[b] = mutableListOf()
            }
            edge[b]!!.add(a)
        }

        fun small(): Int {
            return edge.keys.size
        }

        fun isGood(i: Int): Boolean {
            val vis = hashSetOf<Int>()
            val que = ArrayDeque<Int>()
            que.add(i)
            vis.add(i)
            while(que.isNotEmpty()) {
                val j = que.removeFirst()
                if(j !in edge) continue
                edge[j]!!.forEach({k ->
                    if(k !in vis){
                        que.add(k)
                        vis.add(k)
                        if(vis.size > K) return true
                    }
                })
            }
            return vis.size > K
        }
        fun large(): Int {
            return (1..N).count { isGood(it) }
        }
//        val res = small()
        val res = large()
        println("Case #${it}: $res")
    }
}
