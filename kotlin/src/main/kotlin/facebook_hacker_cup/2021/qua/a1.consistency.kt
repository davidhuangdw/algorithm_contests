
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val V = "AEIOU".toHashSet()
    val T = readLine()!!.toInt()
    (1..T).forEach {
        val S = readLine()!!
        var (vc, vmax, cc, cmax) = listOf(0, 0, 0, 0)
        var count = hashMapOf<Char, Int>()
        for(ch in S){
            count[ch] = count.getOrDefault(ch, 0) + 1
            if(ch in V){
                vc += 1
                vmax = maxOf(vmax, count[ch]!!)
            }else{
                cc += 1
                cmax = maxOf(cmax, count[ch]!!)
            }
        }
        val min_cost = minOf(vc + (cc-cmax) * 2, cc + (vc - vmax)*2)
        println("Case #${it}: $min_cost")
    }
}
