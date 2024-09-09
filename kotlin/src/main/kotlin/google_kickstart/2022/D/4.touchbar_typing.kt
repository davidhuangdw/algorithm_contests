package google_kickstart.`2022`.D

import kotlin.math.absoluteValue

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val (T, ) = input_ints()
    (1..T).forEach {
        val (N, ) =input_ints()
        val S = input_ints()
        val (M, ) = input_ints()
        val K = input_ints()

        fun small(): Int {
            val ind = hashMapOf<Int, Int>()
            for((i, v) in K.withIndex()){
                ind[v] = i
            }
            return S.map{ind[it]!!}.zipWithNext().sumOf{ (a, b) -> (a-b).absoluteValue }
        }

        fun large(): Int {
            val ind = hashMapOf<Int, MutableList<Int>>()
            for((i, v) in K.withIndex()){
                if(v !in ind){ind[v] = mutableListOf()}
                ind[v]!!.add(i)
            }


            val s = mutableListOf(S[0])
            S.zipWithNext { a, b -> if(b != a){ s.add(b)} }
            val n = s.size;

            val near = hashMapOf<Pair<Int, Int>, List<Int>>()
            for(v in s){
                val pos = ind[v]!!
                pos.zipWithNext().forEach {(l, r) ->
                    (l+1 until r).forEach{i -> near[v to i] = listOf(l, r)}
                }
                (0 until pos[0]).forEach{i -> near[v to i] = listOf(pos[0])}
                (pos.last()+1 until M).forEach{i -> near[v to i] = listOf(pos.last())}
            }

            val cache = mutableMapOf<Pair<Int, Int>, Int>()
            fun dp(k: Int, i:Int): Int{
                if(i >= n) return 0
                if(k to i !in cache){
//                    val mn = ind[s[i]]!!.minOf { j -> (k-j).absoluteValue + dp(j, i+1) } // mid
                    val mn = near[s[i] to k]!!.minOf { j -> (k-j).absoluteValue + dp(j, i+1) } // large
                    cache[k to i] = mn
                }
                return cache[k to i]!!
            }
            val res = ind[s[0]]!!.minOf { j ->
                dp(j, 1)
            }

            return res
        }

//        val res = small()
        val res = large()

        println("Case #${it}: $res")
    }
}
