package google_code_jam.`2017`.`1b`
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val (T, ) = input_ints()
    (1..T).forEach {
        val line = input_ints()
        val (N, R, O, Y, G) = line
        val (B, V) = line.subList(5, 7)
        val ord = "ROYGBV"
        val cnt = ord.withIndex().map{(i, ch) ->
            ch to line[i+1]
        }.toMap() as HashMap

        fun set1(): String{
            val ring = MutableList(N){' '}
            val cs = ord.toList().sortedByDescending{cnt[it]}
            var i = 0
            for(ch in cs){
                for(k in 0 until cnt[ch]!!){
                    if(ring[i] != ' ')
                        i ++
                    ring[i] = ch
                    i = (i+2) % N
                }
            }
            for(i in 0 until N)
                if(ring[i] == ring[(i+1) % N])
                    return "IMPOSSIBLE"
            return ring.joinToString("")
        }

        fun set2(): String {
            val singleOther = hashMapOf('O' to 'B', 'G' to 'R', 'V' to 'Y')
            val IMP = "IMPOSSIBLE"
            for((pair, other) in singleOther){
                val co = cnt[other]!!
                val cp = cnt[pair]!!
                if(cp == 0)
                    continue // bug: missing the case
                if(co - cp > 0){
                    cnt[other] = co - cp
                }else if(co == cp && co+cp == N){
                    return "$pair$other".repeat(co)
                }else return IMP
            }
            val cs = "RYB".toList().sortedByDescending{cnt[it]!!}
            val pos = hashMapOf<Char, Int>()
            val n = cs.sumOf{cnt[it]!!}

            val ring = MutableList(n){" "}
            var i = 0
            for(ch in cs){
                for(k in 0 until cnt[ch]!!){
                    if(ring[i] != " ")
                        i ++
                    if(k == 0){
                        pos[ch] = i
                    }
                    ring[i] = ch.toString()
                    i = (i+2) % n
                }
            }
            for(i in 0 until n)
                if(ring[i] == ring[(i+1) % n])
                    return IMP

            for(ch in "OGV"){
                if(cnt[ch]!! > 0){
                    val other = singleOther[ch]!!
                    val i = pos[other]!!
                    ring[i] = "$other${"$ch$other".repeat(cnt[ch]!!)}"
                }
            }
//            println(ring)
            return ring.joinToString("")
        }
        val r = set2()
        println("Case #${it}: $r")
    }
}
