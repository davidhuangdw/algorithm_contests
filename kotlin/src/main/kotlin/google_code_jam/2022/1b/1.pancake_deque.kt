package google_code_jam.`2022`.`1b`
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, ) = input_ints()
        val D = input_ints()
        var (l, r) = 0 to N-1
        var cur = 0
        var cnt = 0
        while(l <= r){
            val v = minOf(D[l], D[r])
            if(D[l] <= D[r]){
                l ++
            }else{
                r --
            }
            if(v >= cur){
                cnt ++
                cur = v
            }
        }
        println("Case #${it}: $cnt")
    }
}
