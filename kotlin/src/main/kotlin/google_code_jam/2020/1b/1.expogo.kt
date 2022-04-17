package google_code_jam.`2020`.`1b`

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val T = readLine()!!.toInt()
    (1..T).forEach {
        var (x, y) = input_longs()
        var impossible = false
        val path = mutableListOf<Char>()
        val xd = if(x > 0){
            "WE"
        }else{
            x = -x
            "EW"
        }
        val yd = if(y > 0){
            "SN"
        }else{
            y = -y
            "NS"
        }
        while(x > 0 || y > 0){
            if((x+y) % 2 == 0L){
                impossible = true
                break
            }
            val add = if((x+y) == 1L || (x/2 + y/2) % 2 == 1L) -1 else 1
            var dir = ' '
            if(x % 2 == 1L){
                x += add
                dir = if(add > 0) xd[0] else xd[1]
            }else{
                y += add
                dir = if(add > 0) yd[0] else yd[1]
            }
            path.add(dir)
            x /= 2
            y /= 2
        }
        val res = if(impossible) "IMPOSSIBLE" else path.joinToString("")
        println("Case #${it}: $res")
    }
}
