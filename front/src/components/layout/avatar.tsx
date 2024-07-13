import { User, Settings, LogOut } from 'lucide-react';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  // DropdownMenuShortcut,
  DropdownMenuTrigger
} from '@/components/ui/dropdown-menu';
import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar';

export const UserAvatar = () => {
  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Avatar className='h-8 w-8 cursor-pointer'>
          <AvatarImage src='https://github.com/shadcn.png' />
          <AvatarFallback>CN</AvatarFallback>
        </Avatar>
      </DropdownMenuTrigger>
      <DropdownMenuContent className='w-[13rem]'>
        <DropdownMenuLabel>Minha Conta</DropdownMenuLabel>
        <DropdownMenuSeparator />
        <DropdownMenuItem className='cursor-pointer'>
          <User className='mr-2 h-4 w-4' />
          <span>Perfil</span>
          {/* <DropdownMenuShortcut>⇧⌘P</DropdownMenuShortcut> */}
        </DropdownMenuItem>
        <DropdownMenuItem className='cursor-pointer'>
          <Settings className='mr-2 h-4 w-4' />
          <span>Configurações</span>
          {/* <DropdownMenuShortcut>⌘S</DropdownMenuShortcut> */}
        </DropdownMenuItem>
        <DropdownMenuSeparator />
        <DropdownMenuItem className='cursor-pointer'>
          <LogOut className='mr-2 h-4 w-4' />
          <span>Sair</span>
          {/* <DropdownMenuShortcut>⇧⌘Q</DropdownMenuShortcut> */}
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
};
