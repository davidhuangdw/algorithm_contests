package google_code_jam.`2018`.r2
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val (T, ) = input_ints()
    (1..T).forEach {
        val (N, ) = input_ints()
        val A = (1..N).map{input_ints()}
        fun set2(): Int{
            val cnt = hashMapOf<Int, MutableList<Pair<Int, Int>>>()
            for((i, row) in A.withIndex()){
                for((j, v) in row.withIndex()){
                    if(v !in cnt)
                        cnt[v] = mutableListOf()
                    cnt[v]!!.add(i to j)
                }
            }

            fun maxMatch(all: List<Pair<Int, Int>>): Int{
                val edge = hashMapOf<Int, MutableList<Int>>()
                for((a, b) in all){
                    if(a !in edge)
                        edge[a] = mutableListOf()
                    edge[a]!!.add(b)
                }

                val bm = hashMapOf<Int, Int>()
                val vis = hashSetOf<Int>()
                fun dfs(a: Int): Boolean{
                    for(b in edge[a]!!){
                        if(b in vis)
                            continue
                        vis.add(b)
                        if(b !in bm || dfs(bm[b]!!)) {
                            bm[b] = a
                            return true
                        }
                    }
                    return false
                }
                var max = 0
                for(a in edge.keys){
                    vis.clear()
                    if(dfs(a))
                        max ++
                }
                return max
            }
            return N*N  - cnt.values.sumOf{maxMatch(it)}
        }
        val res = set2()
        println("Case #${it}: $res")
    }
}
